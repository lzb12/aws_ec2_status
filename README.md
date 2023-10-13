本地开发
环境要求
go 1.20.3


dockerfile

docker build -t aws_ec2_status:1.0 .

docker run --name aws_ec2_status  aws_ec2_status:1.0 -r "ap-northeast-1" -corpid "" -corpsecret "" -touser "" -agentid 100 -awsid "" -aws_secret ""