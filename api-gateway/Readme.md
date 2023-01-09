<!-- PROJECT LOGO -->
<br />
<div align="center">

  <h1>:)</h1>

<h3 align="center">Api Gateway</h3>

  <p align="center">
    Layer which is the only gateway for the client. At this layer, the API Gateway can handle requests from clients and as a security layer, checks whether each request from the client is allowed to proceed or not.
    <br />
    <a href="https://github.com/github_username/repo_name"><strong>Explore the docs »</strong></a>
  </p>
</div>





<!-- ABOUT THE PROJECT -->
## About The Project

<p align="center">
  <img src="https://i.ibb.co/vxbSY2m/api-gateway.jpg" alt="A4-1" border="0" />
</p>

on this api gateway there is a process that will run. after the client accesses the api gateway, it will hit the specified endpoint/route. then the system will check whether the controller at the endpoint must go through the `middleware` or not. after checking the program will access the `service` intended by the controller.

<p align="right">(<a href="#readme-top">back to top</a>)</p>



### Built With

* ![NodeJS](https://img.shields.io/badge/node.js-6DA55F?style=for-the-badge&logo=node.js&logoColor=white)
* ![Express.js](https://img.shields.io/badge/express.js-%23404d59.svg?style=for-the-badge&logo=express&logoColor=%2361DAFB)
*	![JWT](https://img.shields.io/badge/JWT-black?style=for-the-badge&logo=JSON%20web%20tokens)

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- GETTING STARTED -->
## Getting Started


### Prerequisites

you have to install node.js and golang on your machine to run this project. the way to run the service will be in each service folder


### Installation

1. Download and Install Node.JS at [https://nodejs.org](https://nodejs.org/en/download/)
3. Setup service url 
   ```sh
   URL_SERVICE_MEDIA = "http://localhost:3000"
   URL_SERVICE_USERS = "http://localhost:3001"
   URL_SERVICE_COURSE = "http://localhost:3002"
   URL_SERVICE_TRANSACTION = "http://localhost:3003"
   ```
2. Install NPM packages
   ```sh
   npm install
   npm run start
   ```

<p align="right">(<a href="#readme-top">back to top</a>)</p>


<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- ROADMAP -->
## Integrate Roadmap

- [X] Media Service
- [X] User Service
- [X] Course Service
- [ ] Transaction Service

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- LICENSE -->
## License

Distributed under the MIT License.

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- CONTACT -->
## Contact

Kawe - [@kawe123_](https://www.instagram.com/kawe123_/) - muhammad.ilham.kusumawardhana@gmail.com

Project Link: [https://github.com/Ilhamkawe/Simple-Microservice-project/](https://github.com/Ilhamkawe/Simple-Microservice-project/)

<p align="right">(<a href="#readme-top">back to top</a>)</p>
