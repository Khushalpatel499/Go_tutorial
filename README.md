# Go_tutorial
1.It is complied:a programming language that is converted into machine code so that the processor can execute it.
2.go tool run file directy without any virtual machine.
3.create the module by: go mod init hello.
4. for run program : go run main.go
5.lexer(check the grammer is properly used) automatically put semilcolumn so not need to write which is hidden.
6. case sensitive almost and everything is a type.
7. Basic types: string,bool,integer,floating, complex.
8. advanced type:array,slices,maps,structs,pointers,function,channel.
9. TO create a variable used var.
10.First letter capital is mainly we consider as public.
11.for reading we used two library bufio and os
12. bufio.NewReader(os.Stdin)
13. comma ok || err err for handle err or input(input, _)
14. for input : input, _ and for error: _ ,err
15. strconv. (conver string into number)
16. panic() : end the program for handle err in code.
17. build .exe file or other system by GOOS="darwin" go bulid 
18.Memory allocation and deallocation is happen automatically
