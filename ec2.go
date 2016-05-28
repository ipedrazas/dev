package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/awslabs/aws-sdk-go/service/ec2"
	"net/url"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func awsUp() {
	svc := ec2.New(session.New())

	runResult, err := svc.RunInstances(&ec2.RunInstancesInput{
		// An Amazon Linux AMI ID for t2.micro instances in the us-west-2 region
		ImageId:      aws.String("ami-b19e0fc2"),
		InstanceType: aws.String("t2.micro"),
		MinCount:     aws.Int64(1),
		MaxCount:     aws.Int64(1),
	})

	params := &ec2.RunInstancesInput{
		ImageId:  aws.String("ami-b19e0fc2"), // Required
		MaxCount: aws.Int64(1),               // Required
		MinCount: aws.Int64(1),               // Required
		BlockDeviceMappings: []*ec2.BlockDeviceMapping{
			{ // Required
				DeviceName: aws.String("/dev/sda1"),
				Ebs: &ec2.EbsBlockDevice{
					DeleteOnTermination: aws.Bool(true),
					Encrypted:           aws.Bool(true),
					Iops:                aws.Int64(1),
					SnapshotId:          aws.String("snap-506d62bc"),
					VolumeSize:          aws.Int64(10),
					VolumeType:          aws.String("GP2"),
				},
				NoDevice:    aws.String("String"),
				VirtualName: aws.String("String"),
			},
			// More values...
		},

		DisableApiTermination: aws.Bool(true),
		DryRun:                aws.Bool(true),
		EbsOptimized:          aws.Bool(true),

		InstanceInitiatedShutdownBehavior: aws.String("ShutdownBehavior"),
		InstanceType:                      aws.String("t2.nano"),
		Monitoring: &ec2.RunInstancesMonitoringEnabled{
			Enabled: aws.Bool(false), // Required
		},
		NetworkInterfaces: []*ec2.InstanceNetworkInterfaceSpecification{
			{ // Required
				AssociatePublicIpAddress: aws.Bool(true),
				DeleteOnTermination:      aws.Bool(true),
				Description:              aws.String("String"),
				DeviceIndex:              aws.Int64(1),
			},
			// More values...
		},
		SecurityGroupIds: []*string{
			aws.String("sg-f7c66d90"), // Required
			// More values...
		},
	}
	resp, err := svc.RunInstances(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func awsDown() {

	svc := ec2.New(session.New())

	params := &ec2.TerminateInstancesInput{
		InstanceIds: []*string{ // Required
			aws.String("i-8a743c00"), // Required
			// More values...
		},
		DryRun: aws.Bool(false),
	}
	resp, err := svc.TerminateInstances(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func awsList() {

	svc := ec2.New(session.New())

	request := ec2.DescribeInstancesInput{
		Filters: []*ec2.Filter{
			&ec2.Filter{
				Name: aws.String("instance-state-name"),
				Values: []*string{
					aws.String("running"),
					aws.String("pending"),
				},
			},
		}}
	result, err := svc.DescribeInstances(&request)
	check(err)

	// Loop through the instances. They don't always have a name-tag so set it
	// to None if we can't find anything.
	for idx, _ := range result.Reservations {
		for _, inst := range result.Reservations[idx].Instances {

			// We need to see if the Name is one of the tags. It's not always
			// present and not required in Ec2.
			name := "None"
			for _, keys := range inst.Tags {
				if *keys.Key == "Name" {
					name = url.QueryEscape(*keys.Value)
				}
			}

			important_vals := []*string{
				inst.InstanceId,
				&name,
				inst.PrivateIpAddress,
				inst.InstanceType,
				inst.PublicIpAddress,
			}

			// Convert any nil value to a printable string in case it doesn't
			// doesn't exist, which is the case with certain values
			output_vals := []string{}
			for _, val := range important_vals {
				if val != nil {
					output_vals = append(output_vals, *val)
				} else {
					output_vals = append(output_vals, "None")
				}
			}
			// The values that we care about, in the order we want to print them
			fmt.Println(strings.Join(output_vals, " "))
		}
	}
}
