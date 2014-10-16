## A tutorial on creating Go Web App.

### Docker based installation (and running)

```
docker run --publish 8080:8080 maddyonline/gowebapp
```

Visit `http://localhost:8080/hi` to see the app in action. (Use docker VM's IP address instead of localhost if docker is running in VM as in Mac OSX.)


### More Elaborate Instructions

```
# Start
sudo docker run -d --publish 8080:8080 --name mygowebapp  maddyonline/gowebapp

# Stop
sudo docker stop mygowebapp 

# Delete
sudo docker rm mygowebapp

# Restart
sudo docker run -d --publish 8080:8080 --name mygowebapp  maddyonline/gowebapp
```
