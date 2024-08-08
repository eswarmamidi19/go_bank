package main

func main(){
   server := NewApiServer(":5000");
   server.Run();
}
