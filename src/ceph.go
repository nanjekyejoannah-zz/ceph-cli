package ceph

import(

	"flag"
    "fmt"
    "os"
    "cephclient"
)

func main() {
    // subcommands
    connectCommand := flag.NewFlagSet("connect", flag.ExitOnError)
    createbucketCommand := flag.NewFlagSet("createbucket", flag.ExitOnError)
    writefileCommand := flag.NewFlagSet("writefile", flag.ExitOnError)

    //connect subcommand pointers
    connectHostPtr := connectCommand.String("host", "", "host. (Required)")
    connectPortPtr := connectCommand.String("port", "", "port. (Required)")

    //createbucket subcommand pointers
    createbucketbucketnamePtr := createbucketCommand.String("bucketname", "", "bucketname. (Required)")

    //writefile subcommand pointers
    writefilebucketnamePtr := writefileCommand.String("bucketname", "", "bucketname. (Required)")
    writefilefilepathPtr := writefileCommand.String("filepath", "", "filepath. (Required)")

    if len(os.Args) < 2 {
        fmt.Println("subcommand is required")
        os.Exit(1)
    }

    switch os.Args[1] {
    case "connect":
        connectCommand.Parse(os.Args[2:])
    case "createbucket":
        createbucketCommand.Parse(os.Args[2:])
    case "writefile":
        writefileCommand.Parse(os.Args[2:])
    default:
        flag.PrintDefaults()
        os.Exit(1)
    }

    if connectCommand.Parsed() {
        
        if *connectHostPtr == "" {
            connectCommand.PrintDefaults()
            os.Exit(1)
        }

        _, err = cephclient.NewCephClient (*connectHostPtr, *connectPortPtr)

        if err == nil{
        	fmt.Println("connected to rados on %s at port %s", *connectHostPtr, *connectPortPtr)
        }
        else{
        	fmt.Println("Error connecting to rados on %s at port %s", *connectHostPtr, *connectPortPtr)
        }
        
    }



}

