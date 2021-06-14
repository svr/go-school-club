# go-school-club
## Test task for Go school

Docker image for the app was published to [Docker hub](https://hub.docker.com/r/svr3/go-school-club)

To run the app on a local machine the following command can be used:

```
docker container run --rm -p 3000:3000 svr3/go-school-club:0.0.3
```
and point browser to [http://localhost:3000](http://localhost:3000)

The app was also deployed to AWS ECS: [http://go-school-alb-1771026626.eu-central-1.elb.amazonaws.com](http://go-school-alb-1771026626.eu-central-1.elb.amazonaws.com/)
