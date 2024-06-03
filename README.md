# Dataverse Ledger

![image](https://github.com/Dataverse-Ledger/Dataverse-Ledger/assets/97180942/eafc7f88-3f5d-4906-a052-9cd5882fac53)

## Run 

```./scripts/run.sh```
```
./build/dataverse-cli key import demo.pk
./build/dataverse-cli chain import-anr
```

```./build/dataverse-cli chain watch```

```./build/dataverse-cli server start```

## Description
Dataverse is a cutting-edge blockchain platform designed to revolutionize the landscape of data trading and sharing. In today’s digital age, where data integrity and security are paramount, Dataverse stands out as a beacon of trust and reliability. Leveraging the power of blockchain technology, Dataverse ensures that every data transaction is secure, transparent, and verifiable.

At its core, Dataverse serves as a secure registry for devices and sensors, each assigned a unique identifier for seamless data authentication. This innovative platform notarizes data generated by these registered devices, safeguarding it against tampering or unauthorized access. Through Dataverse, users can confidently trade and share data, knowing that its integrity remains intact throughout every transaction.

Key features of Dataverse include device and sensor registration, data notarization, verification, and secure data trading. Whether you're a researcher looking to securely share and monetize your data, a business seeking to leverage IoT insights, or an individual concerned about data privacy, Dataverse provides a robust solution for your data integrity and trust needs.

Join us in embracing a future where data trading is transparent, secure, and accessible to all, powered by Dataverse.

![image](https://github.com/Dataverse-Ledger/Dataverse-Ledger/assets/68425016/b4534a15-9602-4b02-b407-d5c35f935c8e)

## What it does
Dataverse Ledger is a blockchain-based ecosystem designed for secure IoT data trading. It uses a fast and efficient Virtual Machine (VM) and the Snowman consensus algorithm from the Avalanche Hypersdk framework to notarize data autonomously on-chain. This ensures data authenticity and enables producers to trade data securely with native tokens, while allowing other parties to easily verify the data.

## How we built it
We built Dataverse Ledger using the Avalanche Hypersdk framework to create a custom VM, enhancing scalability and transaction speed. Additionally, we developed a custom software/firmware for smart IoT devices, modifying a fork of Tasmota for seamless integration with Dataverse Ledger. This firmware allows devices to autonomously notarize data, ensuring security and authenticity without human intervention.

