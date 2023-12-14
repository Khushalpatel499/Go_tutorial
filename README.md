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
19. There are two methods :1.new() and 2. make()
20. new():  1.Allocate memory but not init
            2.you will get a memory address
            3.zeroed storage(can't put any data initially)
21. make(): 1.Allocate memory and init
            2.you will get a memory address
            3.non-zeroed storage(put any data )
22. GC(garbage collection) happen automatically
23. some time we passed on variables the copy is passed on ,so we want to passed on actual values so pointers are used.
24. array with more features is called slices.
25.append is used to add data in slices
26.no inheritance in golang
27.Function inside function not allowed
28.when function are used in struct called methods
29. The defer keyword is basically whatever we execute is execute line by line but by use of defer that line execute at the very end.They are lifo(last in first out)
30. panic() sutdown the excecution and show what error is coming
31. http is fastest to handle webrequest in go lang.
32. & is like a comma in url world.
33.interface alternate of struct.
34. omitempty if field is empty don't through the field
35.gorilla mux a routing system. command : go get -u github.com/gorilla/mux


