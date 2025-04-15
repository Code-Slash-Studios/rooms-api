The purpose of this repository is to run the API application side of CIS Rooms.

The API backend must be connected to this application so that communication between the database, web application, and Raspberry Pi can flow.

Within the directory's root are files that run the main code for the API.

Under src, authentication is used for the authentication import, controllers is used for the controllers import, models is used for struct definitions, router is used for the routing subsystem, and routes is used to define the various routes for the router.

***NOTE***

Within main.go, there is a database connection that needs to be updated to your environment, including the IP address for the database and user/password information.

We also recommend setting up your environment similar to our environment for easy use. That means allowing routing from port 6000 to 80/api and allowing routing to port 80/pi. See the API code for more details.
