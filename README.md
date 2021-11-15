**Komiser is back 🎉 and rework all of this project !**

<a href="">
  <img src="https://s3.eu-west-3.amazonaws.com/komiser-assets/images/logo.png" width="200" align="right" alt="Amp Logo">
</a>

**Highlights**

* Analyze and manage cloud cost, usage, security, and governance in one place.
* Control your usage and create visibility across all used services to achieve maximum cost-effectiveness.
* Detect potential vulnerabilities that could put your cloud environment at risk.
* Get a deep understanding of how you spend on the AWS, GCP, OVH, DigitalOcean and Azure.


### Homebrew
 ````
 brew install flemzord/tap/komiser
 ````
### Docker:

```
docker run -d -p 3000:3000 -e AWS_ACCESS_KEY_ID="" -e AWS_SECRET_ACCESS_KEY="" -e AWS_DEFAULT_REGION="" --name komiser ghcr.io/flemzord/komiser:latest
```

## How to use

### AWS

* Create an IAM user with the following IAM [policy](https://raw.githubusercontent.com/flemzord/komiser/main/policy.json):

```
wget https://komiser.s3.amazonaws.com/policy.json
```

* Add your **Access Key ID** and **Secret Access Key** to *~/.aws/credentials* using this format

``` 
[default]
aws_access_key_id = <access key id>
aws_secret_access_key = <secret access key>
region = <AWS region>
```

* That should be it. Try out the following from your command prompt to start the server:

```
komiser start --port 3000
```

You can also use Redis as a caching server:

```
komiser start --port 3000 --redis localhost:6379 --duration 30
```

* Point your browser to http://localhost:3000

<p align="center">
    <img src="https://s3.eu-west-3.amazonaws.com/komiser-assets/images/dashboard-aws.png"/>
</p>

#### Multiple AWS Accounts Support

Komiser support multiple AWS accounts through named profiles that are stored in the `config` and `credentials files`. You can configure additional profiles by using `aws configure` with the `--profile` option, or by adding entries to the `config` and `credentials` files.

The following example shows a credentials file with 3 profiles (production, staging & sandbox accounts):

```
[Production]
aws_access_key_id=<AWS_ACCESS_KEY_ID>
aws_secret_access_key=<AWS_SECRET_ACCESS_KEY>

[Staging]
aws_access_key_id=<AWS_ACCESS_KEY_ID>
aws_secret_access_key=<AWS_SECRET_ACCESS_KEY>

[Sandbox]
aws_access_key_id=<AWS_ACCESS_KEY_ID>
aws_secret_access_key=<AWS_SECRET_ACCESS_KEY>
```

To enable multiple AWS accounts feature, add the --multiple option to Komiser:

```
komiser start --port 3000 --redis localhost:6379 --duration 30 --multiple
```

* If you point your browser to http://localhost:3000, you should be able to see your accounts:

<p align="center">
    <img src="https://s3.eu-west-3.amazonaws.com/komiser-assets/images/dashboard-aws-multiple.png"/>
</p>

## Versioning

We use [SemVer](http://semver.org/) for versioning. For the versions available, see the [tags on this repository](https://github.com/flemzord/komiser/tags). 


## License
[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2Fflemzord%2Fkomiser.svg?type=large)](https://app.fossa.io/projects/git%2Bgithub.com%2Fflemzord%2Fkomiser?ref=badge_large)
