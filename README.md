# Desciption:

  This lab push a basic API using docker-compose. 
  
    The infrasrtucture consists of...
    
      (1) Mysql database server
          - It has init scrips to create a user and a database, which we will connect to...
          - It also has an external volume, which will make the data persistant... 
          
      (1) Golang API server, which will interact with the database....
  
  
  
  
# Deployment

   1. Clone the repo to a local folder
      git clone (https://github.com/E-N-G-XOR/Poll_Lap_Api.git)
      
   2. Run deployment command
      docker-compose up --build -d
      
   3. Afer a moments, access the api at: http://localhost:5000
      Should be seeing a nice empty table, the database will be empty, but when we add some data, it will appear here....
      
 
# Using the Api....
    
    Endpoints :
    
       /insert  > Used to add data into DB and accepts a "name" and "timestamps" field. 
       
       /new     >  Used to add data in a online form. It interacts with the /insert.
       
       /index   >  Used to dump the Id, Name, timestamps info in a nice table.... For verification purposes...
       
    Curl statements:
    
    Linux  >
    
        curl -X POST -H 'Content-Type: application/json' -d '{"name":"test user","timestamps":"something"}' http://localhost:5000/insert
      
    Windows  >
    
        $postParams = @{name='poweer';timestamps='test'}
        Invoke-WebRequest -Uri http://localhost:5000/insert -UseBasicParsing -Method POST -Body $postParams
      

 
# Improments:
 
###   Monitoring: 
  
       1. The docker conatiners automatically log to STOUT and STERR, which was can grab with a monitoring stack like EFK.

###   Platform

       1. Hashicorp vualt should be implemented to store the username and passwords...

       2. Input validation should be implemented for the golang frontend to sanitize the inputs...


###   Scaling

       1. queueing should be used to store handle the transaction process.

       2. Scaling in a docker swarm or kubernetes deployment should be used. So we can scale the live systems.... Not included here.

 
