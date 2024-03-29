package controller

import (
	"encoding/hex"
	"github.com/Alan-Lxc/crypto_contest/dcssweb/common"
	"github.com/Alan-Lxc/crypto_contest/dcssweb/dto"
	model "github.com/Alan-Lxc/crypto_contest/dcssweb/model"
	reponse "github.com/Alan-Lxc/crypto_contest/dcssweb/response"
	model1 "github.com/Alan-Lxc/crypto_contest/src/model"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"strconv"
)

var metadatapath = "/home/gary/GolandProjects/crypto_contest4/DCSSmain/src/metadata"

func max(a, b int) int {
	//fmt.Println(a,b)
	if a < b {
		return b
	}
	return a
}
func GetUnitList(ctx *gin.Context) {
	db := common.GetDB()

	secretid, err := strconv.Atoi(ctx.Query("secretid"))
	if err != nil {
		reponse.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "查询错误")
		return
	}
	var secretunits []model1.Secretshare
	var list []dto.UnitListDto
	db.Where("secret_id=? and row_num=0", secretid).Order("unit_id").Find(&secretunits) //

	//fmt.Println(secretunits)
	for i := 0; i < len(secretunits); i++ {
		var newunit = model.Unit{}
		db.Where("unit_id=?", secretunits[i].UnitId).First(&newunit)
		//fmt.Println(len(secretunits[i].Data),secretunits[i].Data)
		tmpString := hex.EncodeToString(secretunits[i].Data[max(0, len(secretunits[i].Data)-2):len(secretunits[i].Data)])
		for len(tmpString) < 4 {
			tmpString = "0" + tmpString
		}
		newUnitListDto := dto.UnitListDto{
			UnitId:             newunit.UnitId,
			UnitIp:             newunit.UnitIp,
			SecretshareContent: tmpString,
		}
		list = append(list, newUnitListDto)
	}
	reponse.Success(
		ctx,
		gin.H{
			"total":    len(list),
			"unitlist": list,
		},
		"获取成功",
	)
}

func GetUnitLog(ctx *gin.Context) {
	db := common.GetDB()

	unitid, err := strconv.Atoi(ctx.Query("unitid"))
	if err != nil {
		reponse.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "查询错误")
		return
	}
	secretid, err := strconv.Atoi(ctx.Query("secretid"))
	if err != nil {
		reponse.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "查询错误")
		return
	}
	var unit model.Unit
	db.Where("unit_id=?", unitid).First(&unit)

	fileName := metadatapath + "/Screct" + strconv.Itoa(int(secretid)) + "Node" + strconv.Itoa(int(unitid)) + ".log"
	logData, err := ioutil.ReadFile(fileName)
	if err != nil {
		//panic(fileName)
	}

	//前20位是位置信息
	//str := "2021/10/15 01:58:42 Node 1 new done,ip = 127.0.0.1:10001\n2021/10/15 01:58:43 [Node 1] Now Get start Phase1\n2021/10/15 01:58:43 [Node 1] send point message to [Node 2]\n2021/10/15 01:58:43 [Node 1] send point message to [Node 3]\n2021/10/15 01:58:43 [Node 1] send point message to [Node 4]\n2021/10/15 01:58:43 [Node 1] send point message to [Node 5]\n2021/10/15 01:58:43 Phase 1 :[Node 1] receive point message from [Node 2]\n2021/10/15 01:58:43 Phase 1 :[Node 1] receive point message from [Node 3]\n2021/10/15 01:58:43 Phase 1 :[Node 1] receive point message from [Node 4]\n2021/10/15 01:58:43 Phase 1 :[Node 1] receive point message from [Node 5]\n2021/10/15 01:58:43 [Node 1] read bulletinboard in phase 1\n2021/10/15 01:58:43 [Node 1] receive zero message from [Node 2] in phase 2\n2021/10/15 01:58:43 [Node 1] send message to [Node 2] in phase 2\n2021/10/15 01:58:43 [Node 1] send message to [Node 3] in phase 2\n2021/10/15 01:58:43 [Node 1] send message to [Node 4] in phase 2\n2021/10/15 01:58:43 [Node 1] send message to [Node 5] in phase 2\n2021/10/15 01:58:43 [Node 1] receive zero message from [Node 3] in phase 2\n2021/10/15 01:58:43 [Node 1] receive zero message from [Node 4] in phase 2\n2021/10/15 01:58:43 [Node 1] receive zero message from [Node 5] in phase 2\n2021/10/15 01:58:43 1 has finish _0ShareSum\n2021/10/15 01:58:43 [Node %!d(MISSING)] start verification in phase 2\n2021/10/15 01:58:43 [Node 1] read bulletinboard in phase 2\n2021/10/15 01:58:43 [Node 1] exponentSum: [919527479152314492520017575367191592683263256082405632185091027972819030997943823490183648180157342586948451807488935955194244950067545680105211023924584, 4488708331611103817365843553572127828114720293967289052761970848310698238665156340739995637022331696132887341554588441642277309627493593404039601398517850]\n2021/10/15 01:58:44 [Node 1] receive point message from [Node 4] in phase3\n2021/10/15 01:58:44 node 1 send point message to node 2 in phase 3\n2021/10/15 01:58:44 node 1 send point message to node 3 in phase 3\n2021/10/15 01:58:44 node 1 send point message to node 4 in phase 3\n2021/10/15 01:58:44 node 1 send point message to node 5 in phase 3\n2021/10/15 01:58:44 [Node 1] receive point message from [Node 5] in phase3\n2021/10/15 01:58:44 [Node 1] receive point message from [Node 2] in phase3\n2021/10/15 01:58:44 [Node 1] receive point message from [Node 3] in phase3\n2021/10/15 01:58:44 [Node 1] has finish sharePhase3\n2021/10/15 01:58:44 [Node 1] write bulletinboard in phase 3\n"
	str := string(logData)
	logs := []dto.UnitLogDto{}
	var flag int
	for i := 0; str[i] != '\n'; i++ {
		flag = i
	}
	item := dto.UnitLogDto{
		Timestamp: str[0:19],
		Content:   str[20:flag],
	}
	logs = append(logs, item)
	for i := 0; i < len(str) && i < len(str); i++ {
		if str[i] == '\n' {
			j := i + 1
			if j < len(str) {
				for j = i + 1; str[j] != '\n' && j < len(str); j++ {
				}
				item := dto.UnitLogDto{
					Timestamp: str[i+1 : i+20],
					Content:   str[i+20 : j],
				}
				logs = append(logs, item)
			}
		}
	}
	reponse.Success(
		ctx,
		gin.H{
			"total": len(logs),
			"logs":  logs,
		},
		"获取成功",
	)

}
