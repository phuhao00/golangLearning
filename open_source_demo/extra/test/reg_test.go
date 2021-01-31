package test
//
//import (
//	"context"
//	"errors"
//	"fmt"
//	"github.com/streadway/amqp"
//	"go.mongodb.org/mongo-driver/bson"
//	"go.mongodb.org/mongo-driver/bson/primitive"
//	"go.mongodb.org/mongo-driver/mongo"
//	"go.mongodb.org/mongo-driver/mongo/options"
//	"os"
//	"os/signal"
//	"regexp"
//	"runtime"
//	"strings"
//	"syscall"
//	"testing"
//	"time"
//)
//
//func TestINNN(t *testing.T) {
//	err := E()
//	if err != nil {
//		buf := make([]byte, 1000)
//		buf = buf[:runtime.Stack(buf, false)]
//		fmt.Println(string(buf))
//	}
//	//t.Log(debug.Stack())
//}
//
//func E() error {
//	//runtime.Caller(1)
//	return errors.New("hh")
//	//return errors.New(fmt.Sprintf("%s, %s", string(debug.Stack()), "jjj"))
//}
//
//func TestController_Run(t *testing.T) {
//	uri := "https://stats.fn.sportradar.com/bet365_service/zh/Asia:Shanghai/gismo/stats_season_fixtures2/77189/1"
//	tmpM := strings.ToValidUTF8(`.`+`bet365_service`+`.`, "")
//	fmt.Println(tmpM)
//	isBet365, _ := regexp.MatchString(tmpM, uri)
//	fmt.Println(isBet365)
//
//}
//func GetMongoClient() (*mongo.Client, error) {
//	uri := fmt.Sprintf("mongodb://%s:%s/%s", "10.0.0.102", "27017", "DataCenter")
//	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
//	if err != nil {
//		return nil, err
//	}
//	return client, nil
//}
//
//func TestMongo(t *testing.T) {
//	ctx := context.Background()
//	client, err := GetMongoClient()
//	if err != nil {
//
//	}
//	err = client.Connect(ctx)
//	if err != nil {
//		panic(err)
//	}
//	ctxP, cancel := context.WithTimeout(ctx, 1*time.Second)
//	defer cancel()
//	collection := client.Database("DataCenter").Collection("data_event_team_incident")
//	//{ "qList": { $elemMatch: { "qid": 1, "reorderFlag": 0} } }
//	//findOptions := options.Find()
//	//findOptions.SetLimit(2)
//	//filter := bson.M{"info": bson.M{"$elemMatch": bson.M{"Id": 0, "ObjectId": 869271}}}
//	objID, _ := primitive.ObjectIDFromHex("5fd0c9c4fdec943d93c12d23")
//	filler2 := bson.M{"_id": objID, "info": bson.M{"$elemMatch": bson.M{"$regex": "tournament_stage"}}}
//
//	result, err := collection.Find(ctxP, filler2)
//	if err != nil {
//
//	}
//	for result.Next(ctxP) {
//		fmt.Println(result.Current.String())
//	}
//
//}
//
////docker run -itd --name rabbitmq  -p 5672:5672 -p 15672:15672 rabbitmq:3-management
////docker run -itd --hostname node1 --name rabbitmq1 -e RABBITMQ_DEFAULT_USER=root -e RABBITMQ_DEFAULT_PASS=root -p 15672:15672 -p 5672:5672 -v /data/rabbitmq:/var/lib/rabbitmq rabbitmq:3.8.3-management
////docker run -itd --hostname my-rabbit --name rabbit_test    -p 15672:15672  -p 5672:5672   rabbitmq-management
////[maoti_craw] #
////host="127.0.0.1"
////port="5672"
////user="guest"
////pwd="guest"
////queuename="maoti_craw"
////kind="topic"
////durable=true
////routingkey.raw="craw.bet365_service.raw"
////routingkey.id="craw.bet365_service.id"
////routingkey.publish="craw.bet365_service.publish"
////routingkey.team="craw.bet365_service.team"
////exchangename="maoti_craw"
////internal=false
////nowait=false
////autoack=true
////deliverymode=1 #// 设置消息是否持久化，1： 非持久化 2：持久化
////routines=1
////autodelete=false
////exclusive=false
////noLocal=true
//var deliveryChMsg = make(<-chan amqp.Delivery)
//
//func TestRabbitMq(t *testing.T) {
//	TaskChannel := utils.NewRabbitMQ("127.0.0.1", "5672", "guest", "guest")
//	_, err := TaskChannel.Conn()
//	if err != nil {
//		panic(err)
//	}
//	channel, err := TaskChannel.NewTopicChannel("maoti_craw")
//	if err != nil {
//		panic(err)
//	}
//	fmt.Println("team consume ....")
//	deliveryChMsg, err := channel.Consume([]string{"craw.bet365_service.team"}, "craw.bet365_service.team", true, nil)
//	if err != nil {
//		fmt.Println(deliveryChMsg)
//	}
//	time.Sleep(time.Second * 2)
//	for msg := range deliveryChMsg {
//		fmt.Println(string(msg.Body))
//	}
//	//go ReceiveDelivery()
//	ProCloseSignal := make(chan os.Signal)
//	signal.Notify(ProCloseSignal, syscall.SIGINT, syscall.SIGKILL, syscall.SIGTERM)
//	<-ProCloseSignal
//
//}
//
////func ReceiveDelivery() {
////	tickerB := time.NewTicker(time.Second)
////	for {
////		select {
////		case data := <-deliveryChMsg:
////forr			fmt.Println(string(data.Body))
////		case <-tickerB.C:
////			time.Sleep(time.Second * 1)
////			fmt.Println("uuu")
////		}
////	}
////}
//
//func TestBson(t *testing.T) {
//	tmp := map[string]string{}
//	finalM := bson.M{}
//	for i, i2 := range tmp {
//		finalM[i] = i2
//	}
//}
