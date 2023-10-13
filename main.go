package main

import (
	awsec2 "aws_ec2_api/AwsEC2"
	"flag"
)

func main() {
	var region string
	var corpid string
	var corpsecret string
	var toUser string
	var agentid int
	flag.StringVar(&region, "r", "ap-northeast-1", "aws region,The default is ap-northeast-1")
	flag.StringVar(&corpid, "corpid", "", "Default enterprise WeChat")
	flag.StringVar(&corpsecret, "corpsecret", "", "Default enterprise WeChat")
	flag.StringVar(&toUser, "touser", "", "Notifiers are separated using |")
	flag.IntVar(&agentid, "agentid", 1, "Default enterprise WeChat")
	flag.StringVar(&awsid, "awsid", "", "aws-id")
	flag.StringVar(&aws_secret, "aws_secret", "", "aws_secret")
	flag.Parse()
	awsec2.GetEc2Status(region, corpid, corpsecret, toUser, agentid)

}
