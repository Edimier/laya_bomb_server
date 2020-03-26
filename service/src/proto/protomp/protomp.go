package protomp
var protoMap map[string]string
func init(){
	protoMap = map[string]string{
		"1000":"rpcdatabase.Prop",
		"rpcdatabase.Prop":"1000",
		"1001":"rpcdatabase.ReqPlayerInfo",
		"rpcdatabase.ReqPlayerInfo":"1001",
		"1002":"rpcdatabase.RespPlayerInfo",
		"rpcdatabase.RespPlayerInfo":"1002",
		"2000":"client.ReqLogin",
		"client.ReqLogin":"2000",
		"2001":"client.RespLogin",
		"client.RespLogin":"2001",
		"2002":"client.Prop",
		"client.Prop":"2002",
		"2003":"client.ReqPlayerInfo",
		"client.ReqPlayerInfo":"2003",
		"2004":"client.RespPlayerInfo",
		"client.RespPlayerInfo":"2004",
		"2005":"client.ReqRank",
		"client.ReqRank":"2005",
		"2006":"client.RankInfo",
		"client.RankInfo":"2006",
		"2007":"client.RespRank",
		"client.RespRank":"2007",
		"2008":"client.ReqEnter",
		"client.ReqEnter":"2008",
		"2009":"client.RespEnter",
		"client.RespEnter":"2009",
		"2010":"client.ReqStartGame",
		"client.ReqStartGame":"2010",
		"2011":"client.RespStartGame",
		"client.RespStartGame":"2011",
		"2012":"client.ReqLeave",
		"client.ReqLeave":"2012",
		"2013":"client.RespLeave",
		"client.RespLeave":"2013",
	}
}
func GetProto(key string) (string, bool){
	if value, ok := protoMap[key];ok{
		return value, true
	} else {
		return "", false
	}
}
