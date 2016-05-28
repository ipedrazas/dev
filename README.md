# devc

Little tool to manage remote dev machines in the cloud

A while ago I decided to move all my computing power to the cloud. Ok, not all, but all the development I've been doing should not be tied to the machine I work with. I started usinginstances in Amazon and Google Cloud Platform and at some point I decided that I would like to improve my workflow.

Starting/stopping machines is not a big deal but it's always nice to have as less friction as possible.

I use amazon spot instances or Google `preemptible` machines, because why do you need to pay more when you can pay less?

Anyway, this tool helps me with that workflow. To start a new machine, you will execute:


		devc up aws

To destroy all your machines you just have to do:

		devc down aws



Before `devc` I had to do things like this:

            gcloud compute   instances create dev01    --zone europe-west1-c
            --preemptible  --machine-type "n1-standard-4"    --subnet "default"
            --tags "dev-box"    --image "/ubuntu-os-cloud/ubuntu-1604-xenial-v20160516a"
            --description "remote development"  --boot-disk-size "100"
            --boot-disk-type "pd-standard"     --boot-disk-device-name "instance-1"


## Setup

The tool assumes that you have your amazon credentials in the `~/.aws/credentials` file, you might need to export the region if you don't have a config file

        [default]
        output = json
        region = eu-west-1

