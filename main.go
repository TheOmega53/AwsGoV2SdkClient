// Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
// SPDX - License - Identifier: Apache - 2.0
// snippet-start:[ec2.go-v2.CreateInstance]
package main

func main() {
	//If you need more EC2 code follow this link:
	//https://github.com/awsdocs/aws-doc-sdk-examples/blob/main/gov2/ec2/common

	//Change the DryRun Flag to false if you want to actually create an Instance.
	//CreateNewInstance(true)

	DescribeInstances()

	// DescribeEndpoint("eu-west-3")

	// MonitorInstance("OFF", "i-0bee8faa65e48b3b1")

	// StopAnInstance("i-0bee8faa65e48b3b1", true)
}
