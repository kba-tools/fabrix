# Hyperledger Fabric Network Management CLI

## :beginner: About

This CLI tool helps manage a Hyperledger Fabric network, including creating new network configurations, selecting existing ones, and managing chaincode deployment. It is built using Go and leverages Cobra for command-line functionality.


## :beginner: Features

- **Configuration Management:** Create new network configurations or select existing ones.
- **Network Setup:** Start a Fabric network, view info,
- **Chaincode Deployment:** Install and manage chaincode on the network.
- **Interactive UI:** Interactive prompts for easy usage



### :notebook: Pre-Requisites
Ensure you have Go installed and set up properly.


### :nut_and_bolt: Installation

```sh
# Clone the repository
git clone https://github.com/yourusername/your-repo.git
cd your-repo

# Build the CLI
go build -o fabric-cli

# Move the binary to a directory in your PATH
sudo mv fabric-cli /usr/local/bin/
```

### :package: Usage

- Open a terminal at any location

```sh
./fabric-cli
```

## :wrench: Development

want to contribute to this project? Make a pull request !!!


### :file_folder: File Structure

```
.
├── cmd
├── pkg
│   ├── configs
│   ├── inputs
│   └── prompts
├── main.co
└── README.md
```


### :fire: Contribution

Your contributions are always welcome and appreciated. Following are the things you can do to contribute to this project.

1.  **Report a bug** <br>
    If you think you have encountered a bug, and I should know about it, feel free to report it [here]() and I will take care of it.

2.  **Request a feature** <br>
    You can also request for a feature [here](), and if it will viable, it will be picked for development.

3.  **Create a pull request** <br>
    It can't get better then this, your pull request will be appreciated by the community. You can get started by picking up any open issues from [here]() and make a pull request.


### :cactus: Branches

I use an agile continuous integration methodology, so the version is frequently updated and development is really fast.

1. **`stage`** is the development branch.

2. **`master`** is the production branch.

3. No other permanent branches should be created in the main repository, you can create feature branches but they should get merged with the master.

**Steps to work with feature branch**

1. To start working on a new feature, create a new branch prefixed with `feat` and followed by feature name. (ie. `feat-FEATURE-NAME`)
2. Once you are done with your changes, you can raise PR.

**Steps to create a pull request**

1. Make a PR to `stage` branch.
2. Comply with the best practices and guidelines e.g. where the PR concerns visual elements it should have an image showing the effect.
3. It must pass all continuous integration checks and get positive reviews.

After this, changes will be merged.

## :camera: Gallery

Pictures of your project.


## :lock: License

Add a license here, or a link to it.
