# disys-mandatory2
# daas, lanc & seho

To start the program: 
1. Clone the repo 
2. Change directory to /disys-mandatory2 (main dir of the repo)  
3. Run "docker compose up" or "docker-compose up" (includes building the image, but can be done manually via "docker compose build"/"docker-compose build")

N.B: If nothing happens (except for when a node has critical section, which lasts for 7 seconds) the system might be deadlocked - please restart in that case

Logs for each node is available within the individual containers via Logfile.txt, 
and can be viewed for example by running "cat Logfile.txt" from the CLI of a container. 

Many commits, handle it
