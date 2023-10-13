package awsec2

import (
	wx "aws_ec2_api/sendwx"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

var svc *ec2.EC2

func GetInstances(awsid, aws_secret, Region string) (*ec2.DescribeInstanceStatusOutput, error) {
	sess, _ := session.NewSession(&aws.Config{
		Credentials: credentials.NewStaticCredentials(awsid, aws_secret, ""),
		Region:      aws.String(Region),
	})
	svc = ec2.New(sess)

	result, err := svc.DescribeInstanceStatus(nil)
	// result, err := svc.DescribeInstances(nil)

	if err != nil {
		return nil, err
	}
	return result, nil
}

func GetEc2Ipaddr(instanceids string) (publicipaddr, privateipaddr string) {

	result, err := svc.DescribeInstances(nil)
	if err != nil {
		fmt.Println(err.Error())
	}
	for _, r := range result.Reservations {
		for _, i := range r.Instances {
			if *i.InstanceId == instanceids {
				// fmt.Println(*i.PublicIpAddress)
				// fmt.Println(*i.PrivateIpAddress)
				return *i.PublicIpAddress, *i.PrivateIpAddress
			}
		}

	}
	return "", ""
}

func GetEc2Status(awsid, aws_secret, Region, corpid, corpsecret, toUser string, agentid int) {
	result, err := GetInstances(awsid, aws_secret, Region)
	if err != nil {
		fmt.Println("Got an error retrieving information about your Amazon EC2 instances:")
		fmt.Println(err.Error())
	}
	for _, r := range result.InstanceStatuses {
		//fmt.Println(r) 打印所有返回信息
		if *r.InstanceState.Name == "running" && Region == "ap-southeast-2" {

			if *r.InstanceStatus.Status != "ok" || *r.SystemStatus.Status != "ok" {
				// fmt.Println(*r.InstanceId)
				_, privateipaddr := GetEc2Ipaddr(*r.InstanceId)
				message := fmt.Sprintf("AWS/EC2检查状态失败\n地区:%s\n主机ID:%s\nip:%s", Region, *r.InstanceId, privateipaddr)
				wx.SendWxMessage(corpid, corpsecret, toUser, message, agentid)
			}
		} else if *r.InstanceState.Name == "running" {
			if *r.InstanceStatus.Status != "ok" || *r.SystemStatus.Status != "ok" {
				// fmt.Println(*r.InstanceId)
				publicipaddr, _ := GetEc2Ipaddr(*r.InstanceId)
				message := fmt.Sprintf("AWS/EC2检查状态失败\n地区:%s\n主机ID:%s\nip:%s", Region, *r.InstanceId, publicipaddr)
				wx.SendWxMessage(corpid, corpsecret, toUser, message, agentid)
			}
		} else {
			fmt.Println(fmt.Sprintf("AWS/EC2停止\n主机:%s", *r.InstanceId))
		}
	}
}
