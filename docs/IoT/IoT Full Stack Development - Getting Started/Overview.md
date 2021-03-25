# Overview of IoT Software Development

* IoT projects involve hardware, firmware, and software
* Thorough testing during development is essential
* Firmware updates helps, but typically annoy the users
* Highest priority is security
* Cost versus performance performance

## Plan Ahead

* User interactions
  * user control
  * system feedback
* Data storage
  * how much data
  * where it is stored
* Automation
  * Event handling
  * process automation

## IoT Naming Convention

* User interfaces
* cloud services
* Servers and database
* Wifi
* Direct connections
* Gateway
* End devices

## Test Early, Test Often

* What steps are required for the provisioning process?
* Do the end devices need to be paired to particular gateway or network?
* Is there a packet sniffer available? What types of signal trace and packet decoding features are included?
* what is the firmware development environment? What is the cost for the license if any?
* what software and firmware licenses are included? What others are required?
* Are the software and firmware costs included in the chip or platform purchase?

## Real World Variables

* Validate range and signal coverage in real-world environments
* Measure energy consumption, battery life, packet loss, throughput
* Use real-world payload sizes and incompressible data
* Enable all encryption, authentication, and security protocols

## Security

Security must be designed into every system from day 1. 

## Raspberry Pi OS installation and Security

* https://raspberrypi.org >> download
* Raspbian
* To image the micro SD card, you can use Etcher https://balena.io
* Recommended the direct installation image (Raspbian)
* Light version of Raspbian (no UI)
* Connect an USB keyboard to the Raspberry board, an HDMI cable for the monitor, and a micro USB for the power
* right after starting up, change the default password to something more secure
* connect a LAN network cable on the ethernet port of the board
* run `ip a` or `ifconfig` and note the ip address
* run `sudo apt-get update` to obtain the latest list of packages
* install a virtual private network client `sudo apt-get upgrade`
* `sudo apt-get install openvpn` to install the VPN client
* set a password `sudo rm /etc/sudoers.d/010_pi-nopasswd`
  
For the VPN server, it can be hosted in the cloud but for testing purposes another Raspberry Pi will do just fine.

