package nodes

import (
	"context"
	"errors"
	"github.com/Alan-Lxc/crypto_contest/dcssweb/common"
	"github.com/Alan-Lxc/crypto_contest/src/model"
	//"fmt"
	"io/ioutil"
	"strings"

	//"fmt"
	"github.com/Alan-Lxc/crypto_contest/src/basic/commitment"
	"github.com/Alan-Lxc/crypto_contest/src/basic/interpolation"
	"github.com/Alan-Lxc/crypto_contest/src/basic/point"
	"github.com/Alan-Lxc/crypto_contest/src/basic/poly"
	"github.com/Nik-U/pbc"
	"github.com/golang/protobuf/proto"
	"math/big"
	//"github.com/Nik-U/pbc"
	pb "github.com/Alan-Lxc/crypto_contest/src/service"
	"github.com/ncw/gmp"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"math/rand"
	"net"
	"os"
	"strconv"
	"sync"
	"time"
)

type Node struct {
	//Label of Node
	label int
	//Total number of Nodes
	counter int
	//Degree of polynomial
	degree int
	//the polynomial was set on Z_p
	p *gmp.Int
	//s0 for reconstruction
	s0 *gmp.Int
	// Rand source
	randState *rand.Rand

	//Polynomial State
	secretShares []*point.Point
	//Polynomial State2
	secretShares2 []*point.Point

	//To store the point(shares) sent from other node
	recPoint []*point.Point
	//To recode the share that have already received
	recCounter *int
	//The poly reconstructed with the shares
	recPoly *poly.Poly
	//Mutex to control
	mutex sync.Mutex

	// Proactivization Polynomial
	proPoly *poly.Poly
	//New Polynomials after phase3
	newPoly *poly.Poly

	// Lagrange Coefficient
	lambda []*gmp.Int
	//	0Shares
	_0Shares     []*gmp.Int
	_0ShareSum   *gmp.Int
	_0ShareCount *int
	//Commitment & Witness in Phase 2
	zeroShareCmt *pbc.Element
	zeroPolyCmt  *pbc.Element
	zeroPolyWit  *pbc.Element
	// Counter for New Secret Shares
	shareCnt *int

	//the other nodes in the committee
	Client []*Node

	////Secret shares of node p(a0,y)
	//secretShare []*point.Point
	//IP_address of node
	IpAddress []string
	//board IP address
	ipOfBoard string
	boardList []string
	//clientconn
	clientConn []*grpc.ClientConn
	//nodeService
	nodeService []pb.NodeServiceClient
	//boardconn
	boardConn *grpc.ClientConn
	//boardService
	boardService pb.BulletinBoardServiceClient
	flag_forweb  bool
	// Commitment and Witness from BulletinBoard
	oldPolyCmt      []*pbc.Element
	zerosumShareCmt []*pbc.Element
	zerosumPolyCmt  []*pbc.Element
	zerosumPolyWit  []*pbc.Element
	midPolyCmt      []*pbc.Element
	newPolyCmt      []*pbc.Element

	// Server
	server *grpc.Server
	// [+] Commitment
	dc  *commitment.DLCommit
	dpc *commitment.DLPolyCommit

	// amt
	amtflag1        int
	amtflag2        int
	amtPoly         []*poly.Poly
	amtPolyCmt      []*pbc.Element
	amtBasicPolyCmt []*pbc.Element
	amtPolyCmtSize  int
	amtStep         int
	// Metrics
	totMsgSize *int
	//存ip_list这个文件的地址
	MetadataPath string
	// Initialize Flag
	iniflag *bool
	//secretid
	secretid    int
	s1          *time.Time
	e1          *time.Time
	s2          *time.Time
	e2          *time.Time
	s3          *time.Time
	e3          *time.Time
	Log         *log.Logger
	tott        int64
	t1          int64
	t2          int64
	t3          int64
	Loglocation string
	finish      bool
}

func (node *Node) GetLabel() int {
	if node != nil {
		return node.label
	} else {
		return 0
	}

}

func (node *Node) CreatAmtPoly(pos, l, r int) {
	if l == r {
		//set poly to l-1
		tmp, _ := poly.NewPoly(1)
		tmp.SetCoeffWithInt(1, 1)
		tmp.GetPtrtoConstant().Neg(gmp.NewInt(int64(l)))
		//fmt.Println(pos)
		//fmt.Println(pos,l,r)
		*node.amtPoly[pos] = tmp
		return
	}
	mid := (l + r) >> 1
	node.CreatAmtPoly(pos<<1, l, mid)
	node.CreatAmtPoly(pos<<1|1, mid+1, r)

	//f pos =f lspos * rspos
	node.amtPoly[pos].Multiply(*node.amtPoly[pos<<1], *node.amtPoly[pos<<1|1])
	return
}
func (node *Node) CreatAmtPolyMod(pos, l, r int) {
	node.amtPoly[pos].Mod(node.p)
	if l == r {
		return
	}
	mid := (l + r) >> 1
	node.CreatAmtPolyMod(pos<<1, l, mid)
	node.CreatAmtPolyMod(pos<<1|1, mid+1, r)
	return
}
func (node *Node) CreatAmtPolyCmt(pos, l, r, step int) {
	node.dpc.Commit(node.amtPolyCmt[step], *node.amtPoly[pos])
	if l == r {
		//fmt.Println(node.amtPoly[pos],node.label)
		node.amtStep = step + 1
		return
	}
	mid := (l + r) >> 1
	if node.label <= mid {
		node.CreatAmtPolyCmt(pos<<1, l, mid, step+1)
	} else {
		node.CreatAmtPolyCmt(pos<<1|1, mid+1, r, step+1)
	}
	return
}
func (node *Node) CreatAmtWitness(pos, l, r int, basicPoly poly.Poly) {

	qPoly, _ := poly.NewPoly(node.counter)
	rPoly, _ := poly.NewPoly(node.counter)
	poly.DivMod(basicPoly, *node.amtPoly[pos], node.p, &qPoly, &rPoly)
	//fmt.Println("q&r is", pos, qPoly, rPoly)
	//tmp1 :=make([]*pbc.Element,1)
	//tmp2 :=node.dpc.NewG1()
	//tmp3 :=make([]*pbc.Element,1)
	//tmp4 :=node.dpc.NewG1()
	//tmp1[0]=node.dpc.NewG1()
	//tmp3[0]=node.dpc.NewG1()
	//qPoly.Mod(node.p)
	//node.amtPoly[pos].Mod(node.p)
	//node.dpc.Commit(tmp1[0], qPoly)
	//node.dpc.Commit(tmp2, rPoly)
	//node.dpc.Commit(tmp3[0], *node.amtPoly[pos])
	//node.dpc.Commit(tmp4, basicPoly)
	//tmp4.Div(tmp4,tmp2)
	//fmt.Println("sdasd ", node.dpc.CalcAmtWitness(tmp4,tmp3,tmp1,gmp.NewInt(0),1))
	////
	//tmp5,_:=poly.NewPoly(0)
	//tmp5.Multiply(qPoly,*node.amtPoly[pos])
	//tmp5.Mod(node.p)
	//tmp6:=node.dpc.NewG1()
	//node.dpc.Commit(tmp6,tmp5)
	//fmt.Println("sdssss ", node.dpc.CalcAmtWitness(tmp6,tmp3,tmp1,gmp.NewInt(0),1))
	//fmt.Println(tmp5,qPoly,rPoly)
	node.dpc.Commit(node.amtBasicPolyCmt[pos], qPoly)
	//fmt.Println("test ''",node.amtBasicPolyCmt[pos], qPoly)
	//fmt.Println(node.amtBasicPolyCmt[pos])
	if l == r {
		//fmt.Println("rPoly is ", l, rPoly)
		//rpoly=phi(l)
		return
	}
	mid := (l + r) >> 1
	node.CreatAmtWitness(pos<<1, l, mid, rPoly)
	node.CreatAmtWitness(pos<<1|1, mid+1, r, rPoly)
	return

}
func (node *Node) GroupAmtWitness(id, size int) []*pbc.Element {
	pos, l, r := 1, 0, size
	step := 0
	tmp := make([]*pbc.Element, 0)
	for l != r {
		tt := node.dpc.NewG1()
		tt.Set(node.amtBasicPolyCmt[pos])
		tmp = append(tmp, tt)
		//fmt.Println("poss ",pos,node.amtBasicPolyCmt[pos],tmp," ?????" ,node.amtBasicPolyCmt)
		step = step + 1
		mid := (l + r) >> 1
		//fmt.Println(pos, l, r, mid, id+1)
		if id+1 <= mid {
			pos = pos << 1
			r = mid
		} else {
			pos = pos<<1 | 1
			l = mid + 1
		}
	}
	tt := node.dpc.NewG1()
	tt.Set(node.amtBasicPolyCmt[pos])
	tmp = append(tmp, tt)
	//fmt.Println(len(tmp))
	return tmp
}
func (node *Node) CreatAmt(ployy poly.Poly, size int) {
	tmpPoly := ployy
	tmpPoly.Mod(node.p)
	//fmt.Println(tmpPoly)
	node.CreatAmtPoly(1, 0, size)
	node.CreatAmtPolyMod(1, 0, size)
	node.CreatAmtPolyCmt(1, 0, size, 0)
	//fmt.Println(node.amtPolyCmt)
	if node.label <= node.degree*2+1 {
		node.CreatAmtWitness(1, 0, size, tmpPoly)
	}
	//fmt.Println("ggggggggggggggg")

	return
}

//Server Handler
func (node *Node) Phase1GetStart(ctx context.Context, msg *pb.StartMsg) (response *pb.ResponseMsg, err error) {

	//id := int(msg.GetSecretid())
	node.Log.Printf("[Node %d] Now Get start Phase1", node.label)
	*node.s1 = time.Now()
	node.CreatAmt(*node.recPoly, node.degree*2+1)
	node.amtflag1 = 1
	tmpPoly, _ := poly.NewPoly(node.degree * 2)
	tmpPoly.ResetTo(*node.recPoly)
	for i := 0; i < node.degree*2+1; i++ {
		x := int32(i + 1)
		y := gmp.NewInt(0)
		//w := node.dpc.NewG1()

		tmpPoly.EvalMod(gmp.NewInt(int64(x)), node.p, y)
		//node.dpc.CreateWitness(w, tmpPoly, gmp.NewInt(int64(x)))
		tmpWitness := node.GroupAmtWitness(i, node.degree*2+1)
		for j := len(tmpWitness); j <= node.amtPolyCmtSize; j++ {
			tmpWitness = append(tmpWitness, node.dpc.NewG1())
		}
		//fmt.Println(node.label, i+1, tmpWitness, node.recPoly)
		//fmt.Println("tttt ",tmpWitness)
		node.secretShares2[i] = point.NewPoint(gmp.NewInt(int64(node.label)), y, tmpWitness)
		//fmt.Println(i+1,label,y,w)
	}
	for i := 0; i < node.degree*2+1; i++ {
		node.secretShares[i] = node.secretShares2[i]
	}
	node.SendMsgToNode(node.secretid)
	return &pb.ResponseMsg{}, nil
}

//client
func (node *Node) SendMsgToNode(secretid int) {
	//if *node.iniflag {
	//	node.NodeConnect()
	//	*node.iniflag = false
	//}

	if node.label > node.degree*2+1 {
		return
	}
	p := point.Point{
		X:       node.secretShares[node.label-1].X,
		Y:       node.secretShares[node.label-1].Y,
		PolyWit: node.secretShares[node.label-1].PolyWit,
	}
	node.mutex.Lock()
	//fmt.Println("hh",len(node.secretShares[node.label-1].PolyWit))
	node.recPoint[p.X.Int32()-1] = &p
	//node.recPoint[*node.recCounter] = &p

	//fmt.Println("111",node.label,node.amtStep)
	*node.recCounter += 1
	flag := *node.recCounter == node.degree*2+1
	node.mutex.Unlock()
	if flag {
		*node.recCounter = 0
		node.ClientReadPhase1()
	}
	var wg sync.WaitGroup
	for i := 0; i < node.degree*2+1; i++ {
		if i != node.label-1 {
			node.Log.Printf("[Node %d] send point message to [Node %d]", node.label, i+1)
			////msg := point.Pointmsg{}
			//msg.SetIndex(node.label)
			//msg.SetPoint(node.secretShares[i])
			//(*node.Client[i]).GetMsgFromNode(msg)
			tmpLength := len(node.secretShares[i].PolyWit)
			bb := make([][]byte, tmpLength)
			for j := 0; j < tmpLength; j++ {
				bb[j] = node.secretShares[i].PolyWit[j].CompressedBytes()
			}
			msg := &pb.PointMsg{
				Index:   int32(node.label),
				X:       node.secretShares[i].X.Bytes(),
				Y:       node.secretShares[i].Y.Bytes(),
				Witness: bb,
				//Witness: nil,
			}
			//fmt.Println(node.secretShares[i].X, i+1, node.secretShares[i].Y, node.secretShares[i].PolyWit)
			wg.Add(1)
			go func(i int, msg *pb.PointMsg) {
				defer wg.Done()
				ctx, cancel := context.WithCancel(context.Background())
				_, err := node.nodeService[i].Phase1ReceiveMsg(ctx, msg)
				if err != nil {
					panic(err)
				}
				defer cancel()
			}(i, msg)
		}
	}
	wg.Wait()
	return
}
func (node *Node) Phase1ReceiveMsg(ctx context.Context, msg *pb.PointMsg) (*pb.ResponseMsg, error) {
	//for node.amtflag1 == 0 {
	//	continue
	//}
	_, err := node.GetMsgFromNode(msg)
	//if err!=nil{
	//	panic(err)
	//}
	return &pb.ResponseMsg{}, err
}

func (node *Node) GetMsgFromNode(pointmsg *pb.PointMsg) (*pb.ResponseMsg, error) {
	*node.totMsgSize = *node.totMsgSize + proto.Size(pointmsg)
	index := pointmsg.GetIndex()
	node.Log.Printf("[Node %d] receive point message from [Node %d] in phase 1", node.label, index)
	x := gmp.NewInt(0)
	x.SetBytes(pointmsg.GetX())
	y := gmp.NewInt(0)
	y.SetBytes(pointmsg.GetY())
	//fmt.Println(node.amtStep)
	//fmt.Println("333",node.label,node.amtStep)
	tmpLength := len(pointmsg.Witness)
	witness := make([]*pbc.Element, tmpLength)
	//fmt.Println(node.label,node.amtStep,len(pointmsg.Witness))
	for i := 0; i < tmpLength; i++ {
		witness[i] = node.dpc.NewG1()
		//fmt.Println(witness[i])
		witness[i].SetCompressedBytes(pointmsg.Witness[i])
		//fmt.Println(i,len(witness))
	}
	//fmt.Println("dddd ",witness)
	p := point.Point{
		X:       x,
		Y:       y,
		PolyWit: witness,
	}
	//fmt.Println(p.X, node.label, p.Y, p.PolyWit)
	//Receive the point and store
	node.mutex.Lock()
	//fmt.Println(node.amtStep,len(p.PolyWit))
	node.recPoint[p.X.Int32()-1] = &p
	*node.recCounter += 1
	flag := (*node.recCounter == node.degree*2+1)
	node.mutex.Unlock()
	if flag {
		*node.recCounter = 0
		node.ClientReadPhase1()
	}
	return &pb.ResponseMsg{}, nil
}

func (node *Node) ClientReadPhase1() {
	//if *node.iniflag {
	//	node.NodeConnect()
	//	*node.iniflag = false
	//}
	node.Log.Printf("[Node %d] read bulletinboard in phase 1", node.label)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	stream, err := node.boardService.ReadPhase1(ctx, &pb.RequestMsg{})
	if err != nil {
		node.Log.Fatalf("client failed to read phase1: %v", err)
	}
	//fmt.Println(node.label,node.oldPolyCmt)
	for i := 0; i < node.degree*2+1; i++ {
		msg, err := stream.Recv()
		*node.totMsgSize = *node.totMsgSize + proto.Size(msg)
		if err != nil {
			node.Log.Fatalf("client failed to receive in read phase1: %v", err)
		}
		index := msg.GetIndex()
		polycmt := msg.GetPolycmt()
		node.oldPolyCmt[index-1].SetCompressedBytes(polycmt)
	}
	//fmt.Println(node.label,node.oldPolyCmt)
	x := make([]*gmp.Int, 0)
	y := make([]*gmp.Int, 0)
	polyCmt := node.dpc.NewG1()
	//time.Sleep(5)
	for i := 0; i <= node.degree; i++ {
		p := node.recPoint[i]
		x = append(x, p.X)
		y = append(y, p.Y)
		polyCmt.Set(node.oldPolyCmt[p.X.Int32()-1])
		//node.log.Println(polyCmt,p.X.Int32(),gmp.NewInt(int64(node.label)), p.Y, p.PolyWit)
		//tmpWitness:=node.CalcAmtWitness(p.PolyWit,p.Y)
		//if !node.dpc.VerifyEval(polyCmt, gmp.NewInt(int64(node.label)), p.Y, tmpWitness) {
		//	//node.log.Println(111)
		//	//fmt.Println(node.label, p.X, "FAIL", polyCmt, p.Y, p.PolyWit)
		//	panic("Reconstruction Verification failed")
		//}
		//fmt.Println("p.Y is",node.label,node.amtPolyCmt,node.amtPoly[1])
		if !node.dpc.CalcAmtWitness(polyCmt, node.amtPolyCmt, p.PolyWit, p.Y, node.amtStep) {
			//node.log.Println(111)
			//fmt.Println(node.label, p.X, "FAIL", polyCmt,node.amtPolyCmt, p.PolyWit)
			panic("Reconstruction Verification failed")
		} else {
			//fmt.Println(node.label, p.X, "FAIL", polyCmt,node.amtPolyCmt, p.PolyWit)
			//fmt.Println("okk!!")
		}

	}
	polyp, _ := interpolation.LagrangeInterpolate(node.degree, x, y, node.p)

	node.recPoly.ResetTo(polyp)
	err = node.Phase1WriteOnBorad()
	if err != nil {
		panic(err)
	}
	return
}

func (node *Node) Phase1WriteOnBorad() error {
	node.Log.Printf("[Node %d] write bulletinboard in phase 1", node.label)
	//fmt.Println(node.label, "poly's len is", node.newPoly.GetDegree(), node.newPoly)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	C := node.dpc.NewG1()
	node.dpc.Commit(C, *node.recPoly)
	//fmt.Println(node.label,C)
	//fmt.Println(node.label,"22")
	msg := &pb.Cmt1Msg{
		Index:   int32(node.label),
		Polycmt: C.CompressedBytes(),
	}
	_, err := node.boardService.WritePhase1(ctx, msg)
	//fmt.Println(node.label,C)
	//log.Printf("finish~~")
	return err
}
func (node *Node) Phase1Verify(ctx context.Context, msg *pb.RequestMsg) (*pb.ResponseMsg, error) {
	node.Log.Printf("[Node %d] start verification in phase 1", node.label)
	node.Phase1Readboard()
	return &pb.ResponseMsg{}, nil
}

func (node *Node) Phase1Readboard() {
	node.Log.Printf("[Node %d] read bulletinboard in phase 1", node.label)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	stream, err := node.boardService.ReadPhase12(ctx, &pb.RequestMsg{})
	if err != nil {
		log.Fatalf("client failed to read phase1: %v", err)
	}
	for i := 0; i < node.degree*2+1; i++ {
		msg, err := stream.Recv()
		*node.totMsgSize = *node.totMsgSize + proto.Size(msg)
		if err != nil {
			log.Fatalf("client failed to receive in read phase1: %v", err)
		}
		index := msg.GetIndex()
		polycmt := msg.GetPolycmt()
		node.oldPolyCmt[index-1].SetCompressedBytes(polycmt)
	}
	*node.e1 = time.Now()
	*node.s2 = time.Now()
	if node.label <= node.degree*2+1 {
		node.ClientSharePhase2()
	}
	return
}

//phase2
func (node *Node) ClientSharePhase2() {
	// Generate Random Numbers
	//fmt.Println("initial 0 is ",node._0Shares[node.counter-1])
	//for i := 0; i < node.counter; i++ {
	//	fmt.Println("lambda ",i+1 ,node.lambda[i])
	//}
	for i := 0; i < node.degree*2+1-1; i++ {
		node._0Shares[i].Rand(node.randState, gmp.NewInt(10))
		inter := gmp.NewInt(0)
		inter.Mul(node._0Shares[i], node.lambda[i])
		node._0Shares[node.degree*2+1-1].Sub(node._0Shares[node.degree*2+1-1], inter)
	}
	//0-share means the S
	node._0Shares[node.degree*2+1-1].Mod(node._0Shares[node.degree*2+1-1], node.p)
	inter := gmp.NewInt(0)
	inter.ModInverse(node.lambda[node.degree*2+1-1], node.p)
	node._0Shares[node.degree*2+1-1].Mul(node._0Shares[node.degree*2+1-1], inter)
	node._0Shares[node.degree*2+1-1].Mod(node._0Shares[node.degree*2+1-1], node.p)
	//////
	//exponentSum := node.dc.NewG1()
	//exponentSum.Set1()
	//Y := make([]*gmp.Int, node.counter)
	//Y[0] = gmp.NewInt(int64(16))
	//Y[1] = gmp.NewInt(int64(24))
	//Y[2] = gmp.NewInt(int64(30))
	//Y[3] = gmp.NewInt(int64(16))
	//Y[4] = gmp.NewInt(int64(0))
	//Y[4].SetString("57896044618658097711785492504343953926634992332820282019728792006155588075461",10)
	////fmt.Println("times ans iss ",exponentSum.Mul(exponentSum,node.dc.NewG1().PowBig(node.zerosumShareCmt[1],big.NewInt(0))))
	//for i := 0; i < node.counter; i++ {
	//	lambda := big.NewInt(0)
	//	lambda.SetString(node.lambda[i].String(), 10)
	//	//lambda.SetString(node.lambda[node.counter-1].String(), 10)
	//	tmp := node.dc.NewG1()
	//	node.dc.Commit(node.zeroShareCmt, Y[i])
	//	fmt.Println(i+1," 's cmt is ",node.zeroShareCmt)
	//	tmp.PowBig(node.zeroShareCmt, lambda)
	//	// log.Printf("label: %d #share %d\nlambda %s\nzeroshareCmt %s\ntmp %s", node.label, i+1, lambda.String(), node.zerosumShareCmt[i].String(), tmp.String())
	//	exponentSum.Mul(exponentSum, tmp)
	//	fmt.Println("exponentsun be ",exponentSum)
	//}
	//fmt.Println("should be 0 but be",exponentSum)
	//for i := 0; i < node.counter; i++ {
	//	tmpppp :=gmp.NewInt(int64(0))
	//	tttmp.EvalMod(gmp.NewInt(int64(i+1)),node.p,tmpppp)
	//	fmt.Println(node.label,tmpppp.Mod(tmpppp.Sub(tmpppp,Y[i]),node.p))
	//}
	//to get sum for \sum_counter
	node.mutex.Lock()
	node._0ShareSum.Add(node._0ShareSum, node._0Shares[node.label-1])
	*node._0ShareCount = *node._0ShareCount + 1
	_0shareSumFinish := (*node._0ShareCount == node.degree*2+1)
	node.mutex.Unlock()

	if _0shareSumFinish {
		node.Log.Printf("[Node %d] has finish _0ShareSum", node.label)
		*node._0ShareCount = 0
		//fmt.Println(node.label,"sum is  ",node._0ShareSum)
		node._0ShareSum.Mod(node._0ShareSum, node.p)
		//fmt.Println(node.label, "sum is  ", node._0ShareSum)

		//get a rand poly_tmp with 0-share
		//rand a poly_tmp polynomial
		node.dc.Commit(node.zeroShareCmt, node._0ShareSum)
		//fmt.Println(node.label," 's sum is ",node._0ShareSum," and the cmt is ",node.zeroShareCmt)
		polyTmp, _ := poly.NewRand(node.degree, node.randState, node.p)
		polyTmp.SetCoeffWithInt(0, 0)
		node.dpc.Commit(node.zeroPolyCmt, polyTmp)
		node.dpc.CreateWitness(node.zeroPolyWit, polyTmp, gmp.NewInt(0))

		err := polyTmp.SetCoeffWithGmp(0, node._0ShareSum)
		if err != nil {
			return
		}
		node.proPoly.ResetTo(polyTmp.DeepCopy())

		node.Phase2Write()
	}
	// share 0-share
	var wg sync.WaitGroup
	for i := 0; i < node.degree*2+1; i++ {
		if i != node.label-1 {
			node.Log.Printf("[Node %d] send message to [Node %d] in phase 2", node.label, i+1)
			msg := &pb.ZeroMsg{
				Index: int32(node.label),
				Share: node._0Shares[i].Bytes(),
			}
			//fmt.Println(node.label,i+1,node._0Shares[i],"send")
			wg.Add(1)
			go func(i int, msg *pb.ZeroMsg) {
				defer wg.Done()
				ctx, cancel := context.WithCancel(context.Background())
				defer cancel()
				_, err := node.nodeService[i].Phase2Share(ctx, msg)
				if err != nil {
					return
				}
			}(i, msg)
		}
	}
	wg.Wait()
	return
}
func (node *Node) Phase2Share(ctx context.Context, msg *pb.ZeroMsg) (*pb.ResponseMsg, error) {
	*node.totMsgSize = *node.totMsgSize + proto.Size(msg)
	index := msg.GetIndex()
	node.Log.Printf("[Node %d] receive zero message from [Node %d] in phase 2", node.label, index)
	inter := gmp.NewInt(0)
	inter.SetBytes(msg.GetShare())

	//fmt.Println(index,node.label,inter,"recive")
	//to get sum for \sum_counter
	node.mutex.Lock()
	node._0ShareSum.Add(node._0ShareSum, inter)
	*node._0ShareCount = *node._0ShareCount + 1
	_0shareSumFinish := (*node._0ShareCount == node.degree*2+1)
	node.mutex.Unlock()

	if _0shareSumFinish {
		node.Log.Printf("[Node %d] has finish _0ShareSum", node.label)
		*node._0ShareCount = 0
		//fmt.Println(node.label,"sum is  ",node._0ShareSum)
		node._0ShareSum.Mod(node._0ShareSum, node.p)
		//fmt.Println(node.label, "sum is  ", node._0ShareSum)

		//get a rand poly_tmp with 0-share
		//rand a poly_tmp polynomial
		node.dc.Commit(node.zeroShareCmt, node._0ShareSum)
		//fmt.Println(node.label," 's sum is ",node._0ShareSum," and the cmt is ",node.zeroShareCmt)
		polyTmp, _ := poly.NewRand(node.degree, node.randState, node.p)
		polyTmp.SetCoeffWithInt(0, 0)
		node.dpc.Commit(node.zeroPolyCmt, polyTmp)
		node.dpc.CreateWitness(node.zeroPolyWit, polyTmp, gmp.NewInt(0))

		polyTmp.SetCoeffWithGmp(0, node._0ShareSum)
		node.proPoly.ResetTo(polyTmp.DeepCopy())

		err := node.Phase2Write()
		if err != nil {
			panic(err)
		}
	}
	return &pb.ResponseMsg{}, nil
}
func (node *Node) Phase2Write() error {
	node.Log.Printf("[Node %d] write bulletinboard in phase 2", node.label)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() // Commitment and Witness from BulletinBoard
	msg := &pb.CommitMsg{
		Index:       int32(node.label),
		ShareCommit: node.zeroShareCmt.CompressedBytes(),
		PolyCommit:  node.zeroPolyCmt.CompressedBytes(),
		ZeroWitness: node.zeroPolyWit.CompressedBytes(),
	}
	//fmt.Println("t1 ",int32(node.label),node._0ShareSum,node.zeroShareCmt)
	_, err := node.boardService.WritePhase2(ctx, msg)
	return err
}
func (node *Node) Phase2Verify(ctx context.Context, request *pb.RequestMsg) (response *pb.ResponseMsg, err error) {
	node.Log.Printf("[Node %d] start verification in phase 2", node.label)
	node.ClientReadPhase2()
	return &pb.ResponseMsg{}, nil
}

func (node *Node) ClientReadPhase2() {
	node.Log.Printf("[Node %d] read bulletinboard in phase 2", node.label)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	stream, err := node.boardService.ReadPhase2(ctx, &pb.RequestMsg{})
	if err != nil {
		node.Log.Fatalf("client failed to read phase2: %v", err)
	}
	for i := 0; i < node.degree*2+1; i++ {
		msg, err := stream.Recv()
		*node.totMsgSize = *node.totMsgSize + proto.Size(msg)
		if err != nil {
			node.Log.Fatalf("client failed to receive in read phase1: %v", err)
		}
		index := msg.GetIndex()
		sharecmt := msg.GetShareCommit()
		polycmt := msg.GetPolyCommit()
		zerowitness := msg.GetZeroWitness()

		//fmt.Println("t2 ", index, node.dpc.NewG1().SetCompressedBytes(sharecmt), node.dpc.NewG1().SetCompressedBytes(polycmt), node.dpc.NewG1().SetCompressedBytes(zerowitness), node.label)
		node.zerosumShareCmt[index-1].SetCompressedBytes(sharecmt)
		inter := node.dpc.NewG1()
		inter.SetString(node.zerosumShareCmt[index-1].String(), 10)

		node.zerosumPolyCmt[index-1].SetCompressedBytes(polycmt)
		node.midPolyCmt[index-1].Mul(inter, node.zerosumPolyCmt[index-1])
		node.zerosumPolyWit[index-1].SetCompressedBytes(zerowitness)

	}
	exponentSum := node.dc.NewG1()
	//fmt.Println(exponentSum)
	exponentSum.Set1()
	//fmt.Println("times ans iss ",exponentSum.Mul(exponentSum,node.dc.NewG1().PowBig(node.zerosumShareCmt[1],big.NewInt(0))))
	for i := 0; i < node.degree*2+1; i++ {
		lambda := big.NewInt(0)
		lambda.SetString(node.lambda[i].String(), 10)
		//lambda.SetString(node.lambda[node.counter-1].String(), 10)
		tmp := node.dc.NewG1()
		tmp.PowBig(node.zerosumShareCmt[i], lambda)
		//node.log.Printf("label: %d #share %d\nlambda %s\nzeroshareCmt %s\ntmp %s", node.label, i+1, lambda.String(), node.zerosumShareCmt[i].String(), tmp.String())
		exponentSum.Mul(exponentSum, tmp)
		//fmt.Println(i+1, " 's cmt is", node.zerosumShareCmt[i], "Hey!!! exponentsun be ", exponentSum)
	}

	flag := true
	for i := 0; i < node.degree*2+1; i++ {
		if !node.dpc.VerifyEval(node.zerosumPolyCmt[i], gmp.NewInt(0), gmp.NewInt(0), node.zerosumPolyWit[i]) {
			flag = false
		}
	}
	if !flag {
		panic("Proactivization Verification 2 failed")
	}
	log.Printf("%d exponentSum: %s", node.label, exponentSum.String())
	if !exponentSum.Is1() {
		panic("Proactivization Verification 1 failed")
	}
	*node.e2 = time.Now()
	*node.s3 = time.Now()
	//fmt.Println("hhh")
	//if node.label <= node.degree*2+1 {
	node.ClientSharePhase3()
	//}
	return
	//}
}

func ReadIpList(metadataPath string) []string {
	ipData, err := ioutil.ReadFile(metadataPath + "/ip_list")
	if err != nil {
		log.Fatalf("node failed to read iplist %v\n", err)
	}
	return strings.Split(string(ipData), "\n")
}

//第三部分的分享， 是全份额重建
//通过新的多项式B（x，i），对包括自己的每一个node发送B（i，i）
//最后重建多项式

func (node *Node) ClientSharePhase3() {
	node.newPoly.Add(*node.recPoly, *node.proPoly)
	node.newPoly.Mod(node.p)
	node.CreatAmt(*node.newPoly, node.counter)
	node.amtflag2 = 1
	if node.label > node.degree*2+1 {
		return
	}
	C := node.dpc.NewG1()
	node.dpc.Commit(C, *node.newPoly)
	//fmt.Println("newpoly is ",node.label,node.newPoly,C)
	//old1 := node.dpc.NewG1()
	//old2 := node.dpc.NewG1()
	//old3 := node.dpc.NewG1()
	//node.dpc.Commit(old1, *node.recPoly)
	//node.dpc.Commit(old2, *node.proPoly)
	//node.dpc.Commit(old3, *node.newPoly)

	//tmp := node.dpc.NewG1()
	//fmt.Println(node.label,"verify 1 ",old1.Equals(node.oldPolyCmt[node.label-1]))
	//fmt.Println(node.label,"verify 2 ",old3.Equals(tmp.Mul(old1,old2)))
	//fmt.Println(node.label," poly is ",*node.recPoly)
	var wg sync.WaitGroup
	for i := 0; i < node.counter; i++ {
		value := gmp.NewInt(0)
		node.newPoly.EvalMod(gmp.NewInt(int64(i+1)), node.p, value)
		//witness := node.dpc.NewG1()
		//node.dpc.CreateWitness(witness, *node.newPoly, gmp.NewInt(int64(i+1)))
		tmpWitness := node.GroupAmtWitness(i, node.counter)
		//fmt.Println(node.label, i+1, tmpWitness, node.newPoly)
		tmpLength := len(tmpWitness)
		bb := make([][]byte, tmpLength)
		for j := 0; j < tmpLength; j++ {
			bb[j] = tmpWitness[j].CompressedBytes()
		}
		if i != node.label-1 {
			node.Log.Printf("[Node %d] send point message to [Node %d] in phase 3", node.label, i+1)

			msg := &pb.PointMsg{
				Index:   int32(node.label),
				X:       gmp.NewInt(int64(node.label)).Bytes(),
				Y:       value.Bytes(),
				Witness: bb,
			}
			//把消息发送给不同的节点
			wg.Add(1)
			go func(i int, msg *pb.PointMsg) {
				ctx, cancel := context.WithCancel(context.Background())
				defer cancel()
				defer wg.Done()
				//fmt.Println(i)
				node.nodeService[i].Phase3SendMsg(ctx, msg)
			}(i, msg)
		} else {
			node.secretShares[i].Y.Set(value)
			node.secretShares[i].PolyWit = tmpWitness
			//fmt.Println("check is" ,len(node.amtPolyCmt),node.amtStep,node.amtPolyCmt,tmpWitness,node.label,node.counter,node.degree,value,
			//	node.dpc.CalcAmtWitness(C,node.amtPolyCmt,tmpWitness,value,node.amtStep))
			node.mutex.Lock()
			*node.shareCnt = *node.shareCnt + 1
			flag := *node.shareCnt == node.degree*2+1
			node.mutex.Unlock()
			if flag {
				node.Log.Printf("[Node %d] has finish sharePhase3", node.label)
				*node.shareCnt = 0
				node.Phase3WriteOnBorad()
			}
		}
	}

	return
}

//重建secretShare
func (node *Node) Phase3SendMsg(ctx context.Context, msg *pb.PointMsg) (*pb.ResponseMsg, error) {
	//for node.amtflag2 == 0 {
	//	continue
	//}
	*node.totMsgSize = *node.totMsgSize + proto.Size(msg)
	index := msg.GetIndex()
	Y := msg.GetY()
	node.Log.Printf("[Node %d] receive point message from [Node %d] in phase3", node.label, index)
	witness := msg.GetWitness()
	//fmt.Println("node index is ",index-1)
	node.secretShares[index-1].Y.SetBytes(Y)
	tmpLength := len(witness)
	//fmt.Println(node.label,node.amtStep,len(pointmsg.Witness))
	for i := 0; i < tmpLength; i++ {
		node.secretShares[index-1].PolyWit[i].SetCompressedBytes(witness[i])
	}
	node.mutex.Lock()
	*node.shareCnt = *node.shareCnt + 1
	flag := *node.shareCnt == node.degree*2+1
	node.mutex.Unlock()
	if flag {
		node.Log.Printf("[Node %d] has finish sharePhase3", node.label)
		*node.shareCnt = 0
		if node.label <= node.degree*2+1 {
			node.Phase3WriteOnBorad()
		}
	}
	return &pb.ResponseMsg{}, nil
}
func (node *Node) Phase3WriteOnBorad() {
	node.Log.Printf("[Node %d] write bulletinboard in phase 3", node.label)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	C := node.dpc.NewG1()
	//fmt.Println(node.label, "poly's len is", node.newPoly.GetDegree(), node.newPoly)
	node.dpc.Commit(C, *node.newPoly)
	//fmt.Println(node.label, "poly's len is", node.newPoly.GetDegree(), node.newPoly)
	//fmt.Println(node.label,C)
	//fmt.Println(node.label,"22")
	msg := &pb.Cmt1Msg{
		Index:   int32(node.label),
		Polycmt: C.CompressedBytes(),
	}
	//err := new(error)
	_, err := node.boardService.WritePhase3(ctx, msg)
	for err != nil {
		//panic(err)
		//panic(err)
	}
	//fmt.Println(node.label,C)
	//log.Printf("finish~~")
	return
}
func (node *Node) Phase3Verify(ctx context.Context, msg *pb.RequestMsg) (*pb.ResponseMsg, error) {
	node.Log.Printf("[Node %d] start verification in phase 3", node.label)
	node.Phase3Readboard()
	return &pb.ResponseMsg{}, nil
}
func (node *Node) Gettot() int {
	return *node.totMsgSize
}
func (node *Node) Gettime() []int64 {
	tt := make([]int64, 4)
	tt[0] = node.e3.Sub(*node.s1).Nanoseconds()
	tt[1] = node.e1.Sub(*node.s1).Nanoseconds()
	tt[2] = node.e2.Sub(*node.s2).Nanoseconds()
	tt[3] = node.e3.Sub(*node.s3).Nanoseconds()
	return tt
}
func (node *Node) Sendtestmsg(ctx context.Context, msg *pb.RequestMsg) (*pb.TestMsg, error) {
	return &pb.TestMsg{
		Totmsgsize: int32(*node.totMsgSize),
		Totaltime:  node.tott,
		Phase1Time: node.t1,
		Phase2Time: node.t2,
		Phase3Time: node.t3,
	}, nil
}
func (node *Node) Phase3Readboard() {
	node.Log.Printf("[Node %d] read bulletinboard in phase 3", node.label)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	stream, err := node.boardService.ReadPhase3(ctx, &pb.RequestMsg{})
	if err != nil {
		node.Log.Fatalf("client failed to read phase3: %v", err)
	}
	for i := 0; i < node.degree*2+1; i++ {
		msg, err := stream.Recv()
		*node.totMsgSize = *node.totMsgSize + proto.Size(msg)
		if err != nil {
			node.Log.Fatalf("client failed to receive in read phase1: %v", err)
		}
		index := msg.GetIndex()
		polycmt := msg.GetPolycmt()
		node.newPolyCmt[index-1].SetCompressedBytes(polycmt)
	}
	//fmt.Println("hhhh")
	tmpX := make([]*gmp.Int, node.degree*2+1)
	tmpY := make([]*gmp.Int, node.degree*2+1)
	//time.Sleep(5)
	for i := 0; i < node.degree*2+1; i++ {
		tmp := node.dpc.NewG1()
		if !node.newPolyCmt[i].Equals(tmp.Mul(node.oldPolyCmt[i], node.midPolyCmt[i])) {
			panic("Share Distribution Verification 1 failed")
		}
		//if !node.dpc.VerifyEval(node.newPolyCmt[i], gmp.NewInt(int64(node.label)), node.secretShares[i].Y, node.secretShares[i].PolyWit) {
		//	fmt.Println(node.newPolyCmt[i], gmp.NewInt(int64(node.label)), node.secretShares[i].Y, node.secretShares[i].PolyWit)
		//	panic("Share Distribution Verification 2 failed")
		//}
		if !node.dpc.CalcAmtWitness(node.newPolyCmt[i], node.amtPolyCmt, node.secretShares[i].PolyWit, node.secretShares[i].Y, node.amtStep) {
			//fmt.Println(node.newPolyCmt[i], node.amtPolyCmt,  node.secretShares[i].PolyWit,node.secretShares[i].Y,)
			panic("Share Distribution Verification 2 failed")
		} else {
			//fmt.Println("ok!!!")
		}

		tmpX[i] = gmp.NewInt(int64(i + 1))
		tmpY[i] = node.secretShares[i].Y
	}

	polyp, _ := interpolation.LagrangeInterpolate(node.degree*2, tmpX, tmpY, node.p)
	//fmt.Println(tmpX,tmpY,node.degree,polyp)
	node.recPoly.ResetTo(polyp)
	//store
	coffes := node.recPoly.GetAllCoeffs()
	tmpLength := len(coffes)
	bb := make([][]byte, tmpLength)
	for j := 0; j < tmpLength; j++ {
		bb[j] = coffes[j].Bytes()
	}
	node.save_secret(node.degree, node.counter, node.secretid, bb)
	//y := gmp.NewInt(0)
	//node.recPoly.EvalMod(gmp.NewInt(int64(0)), node.p, y)
	//fmt.Println(node.label,y)
	*node.e3 = time.Now()
	//f, _ := os.OpenFile(node.MetadataPath+"/log"+strconv.Itoa(node.label), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	//defer f.Close()
	node.Log.Printf("[Node %d] finished handoff", node.label)
	node.Log.Printf("[Node %d] totMsgSize,%d\n", node.label, *node.totMsgSize)
	node.Log.Printf("[Node %d] epochLatency,%d\n", node.label, node.e3.Sub(*node.s1).Nanoseconds())
	node.Log.Printf("[Node %d] reconstructionLatency,%d\n", node.label, node.e1.Sub(*node.s1).Nanoseconds())
	node.Log.Printf("[Node %d] proactivizationLatency,%d\n", node.label, node.e2.Sub(*node.s2).Nanoseconds())
	node.Log.Printf("[Node %d] sharedistLatency,%d\n", node.label, node.e3.Sub(*node.s3).Nanoseconds())
	node.tott = (node.e3.Sub(*node.s1).Nanoseconds())
	node.t1 = (node.e1.Sub(*node.s1).Nanoseconds())
	node.t2 = (node.e2.Sub(*node.s2).Nanoseconds())
	node.t3 = (node.e3.Sub(*node.s3).Nanoseconds())
	//node.Log.Printf("the secret for reconstruction is ,%s\n", node.s0.String())
	//*node.totMsgSize = 0
	for i := 0; i < node.degree*2+1; i++ {
		node._0Shares[i].SetInt64(0)
	}
	node._0ShareSum.SetInt64(0)
	node.amtflag2 = 0
	node.amtflag1 = 0
	node.finish = true
	return
}

func (node *Node) Finish() bool {
	return node.finish
}
func (node *Node) Reconstruct(ctx context.Context, request *pb.RequestMsg) (response *pb.ResponseMsg, err error) {
	node.Log.Printf("[Node %d] start reconstruct ", node.label)
	node.Phase3Readboard2()
	return &pb.ResponseMsg{}, nil
}
func (node *Node) Phase3Readboard2() {
	//node.Log.Printf("[Node %d] write bulletinboard in phase 3 2", node.label)
	//fmt.Println(node.label, "poly's len is", node.newPoly.GetDegree(), node.newPoly)

	node.recPoly.EvalMod(gmp.NewInt(0), node.p, node.s0)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	//C := node.dpc.NewG1()
	//node.dpc.Commit(C, *node.recPoly)
	////fmt.Println(node.label,C)
	////fmt.Println(node.label,"22")
	//msg := &pb.Cmt1Msg{
	//	Index:   int32(node.label),
	//	Polycmt: C.CompressedBytes(),
	//}
	//
	//node.boardService.WritePhase32(ctx, msg)
	tmpPoly, _ := poly.NewPoly(node.degree * 2)
	tmpPoly.ResetTo(*node.recPoly)
	YY := gmp.NewInt(0)
	tmpPoly.EvalMod(gmp.NewInt(0), node.p, YY)
	C2 := node.dpc.NewG1()
	node.dpc.CreateWitness(C2, tmpPoly, gmp.NewInt(0))
	tmpWitness := node.GroupAmtWitness(-1, node.counter)
	tmpLength := len(tmpWitness)
	bb := make([][]byte, tmpLength)
	for j := 0; j < tmpLength; j++ {
		bb[j] = tmpWitness[j].CompressedBytes()
	}
	//fmt.Println(YY)
	msg2 := &pb.PointMsg{
		Index:   int32(node.label),
		X:       gmp.NewInt(int64(node.label)).Bytes(),
		Y:       YY.Bytes(),
		Witness: bb,
	}
	_, err := node.boardService.ReconstructSecret(ctx, msg2)
	if err != nil {
		panic(err)
	}
	return
}

func New(degree int, label int, counter int, logPath string, coeff []*gmp.Int) (Node, error) {
	if label < 0 {
		return Node{}, errors.New("Label must be a non-negative number!")
	}
	//file, _ := os.Create(logPath + "/log" + strconv.Itoa(label))
	//defer file.Close()
	ipRaw := ReadIpList(logPath)[0 : counter+1]
	//fmt.Println(":????")
	bip := ipRaw[0]
	ipList := ipRaw[1 : counter+1]

	if counter < 0 {
		return Node{}, errors.New("Counter must be a non-negative number!")
	}
	randState := rand.New(rand.NewSource(time.Now().Local().UnixNano()))
	//fixedRandState := rand.New(rand.NewSource(int64(3)))
	dc := commitment.DLCommit{}
	dc.SetupFix()
	dpc := commitment.DLPolyCommit{}
	dpc.SetupFix(counter + 1)

	p := gmp.NewInt(0)
	p.SetString("57896044618658097711785492504343953926634992332820282019728792006155588075521", 10)
	lambda := make([]*gmp.Int, counter)
	// Calculate Lagrange Interpolation
	denominator := poly.NewConstant(1)
	tmp, _ := poly.NewPoly(1)
	tmp.SetCoeffWithInt(1, 1)
	//fmt.Println(time.Now())
	for i := 0; i < degree*2+1; i++ {
		tmp.GetPtrtoConstant().Neg(gmp.NewInt(int64(i + 1)))
		denominator.MulSelf(tmp)
	}
	for i := 0; i < degree*2+1; i++ {
		lambda[i] = gmp.NewInt(0)
		deno, _ := poly.NewPoly(0)
		tmp.GetPtrtoConstant().Neg(gmp.NewInt(int64(i + 1)))

		deno.Divide(denominator, tmp)
		deno.EvalMod(gmp.NewInt(0), p, lambda[i])
		inter := gmp.NewInt(0)
		deno.EvalMod(gmp.NewInt(int64(i+1)), p, inter)
		interInv := gmp.NewInt(0)
		interInv.ModInverse(inter, p)
		lambda[i].Mul(lambda[i], interInv)
		lambda[i].Mod(lambda[i], p)
	}
	//fmt.Println(time.Now())

	_0Shares := make([]*gmp.Int, counter)
	for i := 0; i < counter; i++ {
		_0Shares[i] = gmp.NewInt(0)
	}
	_0ShareCount := 0
	_0ShareSum := gmp.NewInt(0)

	recPoint := make([]*point.Point, counter)
	recCounter := 0

	secretShares := make([]*point.Point, counter)
	secretShares2 := make([]*point.Point, counter)
	//tmpPoly, err := poly.NewRand(degree, fixedRandState, p)
	tmpPoly, err := poly.NewPoly(len(coeff) - 1)
	tmpPoly.SetbyCoeff(coeff)
	//fmt.Println(tmpPoly.GetDegree())
	//fmt.Println(tmpPoly)
	if err != nil {
		panic("Error initializing random tmpPoly")
	}

	//fmt.Println(time.Now())
	proPoly, _ := poly.NewPoly(degree)
	recPoly, _ := poly.NewPoly(degree)
	newPoly, _ := poly.NewPoly(degree)
	shareCnt := 0

	s0 := gmp.NewInt(0)
	recPoly.ResetTo(tmpPoly)
	recPoly.EvalMod(gmp.NewInt(0), p, s0)

	oldPolyCmt := make([]*pbc.Element, counter)
	midPolyCmt := make([]*pbc.Element, counter)
	newPolyCmt := make([]*pbc.Element, counter)
	for i := 0; i < counter; i++ {
		oldPolyCmt[i] = dpc.NewG1()
		midPolyCmt[i] = dpc.NewG1()
		newPolyCmt[i] = dpc.NewG1()
	}

	zeroShareCmt := dc.NewG1()
	zeroPolyCmt := dpc.NewG1()
	zeroPolyWit := dpc.NewG1()

	zerosumShareCmt := make([]*pbc.Element, counter)
	zerosumPolyCmt := make([]*pbc.Element, counter)
	zerosumPolyWit := make([]*pbc.Element, counter)

	for i := 0; i < counter; i++ {
		zerosumShareCmt[i] = dc.NewG1()
		zerosumPolyCmt[i] = dpc.NewG1()
		zerosumPolyWit[i] = dpc.NewG1()
	}
	// amt
	amtPoly := make([]*poly.Poly, (counter+5)<<2)
	amtPolyCmt := make([]*pbc.Element, (counter+5)<<2)
	amtBasicPolyCmt := make([]*pbc.Element, (counter+5)<<2)
	amtStep := 0
	amtPolycmtSize := 0
	for (1 << amtPolycmtSize) <= counter+1 {
		amtPolycmtSize += 1
	}
	amtPolycmtSize += 1
	for i := 0; i < ((counter + 5) << 2); i++ {
		tmp2Poly, _ := poly.NewPoly(degree)
		amtPoly[i] = &tmp2Poly
		amtPolyCmt[i] = dpc.NewG1()
		amtBasicPolyCmt[i] = dpc.NewG1()
	}

	totMsgSize := 0
	s1 := time.Now()
	e1 := time.Now()
	s2 := time.Now()
	e2 := time.Now()
	s3 := time.Now()
	e3 := time.Now()

	amtflag1 := 0
	amtflag2 := 0

	clientConn := make([]*grpc.ClientConn, counter)
	nodeService := make([]pb.NodeServiceClient, counter)

	iniflag := true

	fileName := "./src/metadata/logOfNode" + strconv.Itoa(label) + ".log"
	tmplogger, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModePerm)
	if err != nil {
		tmplogger, err = os.Create(fileName)
	}
	os.Truncate(fileName, 0)
	log := log.New(tmplogger, "", log.LstdFlags)
	//log.Printf("[Node %d] new done,ip = %s", label, ipList[label-1])
	return Node{
		flag_forweb:     false,
		MetadataPath:    logPath,
		IpAddress:       ipList,
		ipOfBoard:       bip,
		degree:          degree,
		label:           label,
		counter:         counter,
		randState:       randState,
		dc:              &dc,
		dpc:             &dpc,
		p:               p,
		s0:              s0,
		lambda:          lambda,
		_0Shares:        _0Shares,
		_0ShareCount:    &_0ShareCount,
		_0ShareSum:      _0ShareSum,
		secretShares:    secretShares,
		secretShares2:   secretShares2,
		recPoint:        recPoint,
		recCounter:      &recCounter,
		recPoly:         &recPoly,
		proPoly:         &proPoly,
		newPoly:         &newPoly,
		shareCnt:        &shareCnt,
		oldPolyCmt:      oldPolyCmt,
		midPolyCmt:      midPolyCmt,
		newPolyCmt:      newPolyCmt,
		zeroShareCmt:    zeroShareCmt,
		zeroPolyCmt:     zeroPolyCmt,
		zeroPolyWit:     zeroPolyWit,
		zerosumPolyCmt:  zerosumPolyCmt,
		zerosumPolyWit:  zerosumPolyWit,
		zerosumShareCmt: zerosumShareCmt,
		totMsgSize:      &totMsgSize,
		s1:              &s1,
		e1:              &e1,
		s2:              &s2,
		e2:              &e2,
		s3:              &s3,
		e3:              &e3,
		clientConn:      clientConn,
		nodeService:     nodeService,
		iniflag:         &iniflag,
		Log:             log,
		amtStep:         amtStep,
		amtPolyCmt:      amtPolyCmt,
		amtPoly:         amtPoly,
		amtBasicPolyCmt: amtBasicPolyCmt,
		amtPolyCmtSize:  amtPolycmtSize,
		amtflag1:        amtflag1,
		amtflag2:        amtflag2,
	}, nil

}

func (node *Node) NodeConnect() {
	//fmt.Println(node.label,"connect!!")
	boradConn, err := grpc.Dial(node.ipOfBoard, grpc.WithInsecure())
	if err != nil {
		node.Log.Fatalf("Fail to connect board:%v", err)
	}
	node.boardConn = boradConn
	node.boardService = pb.NewBulletinBoardServiceClient(boradConn)

	for i := 0; i < node.counter; i++ {
		if i != node.label-1 {
			clientconn, err := grpc.Dial(node.IpAddress[i], grpc.WithInsecure())
			if err != nil {
				node.Log.Fatalf("[Node %d] Fail to connect to other node:%v", node.label, err)
			}
			node.clientConn[i] = clientconn
			node.nodeService[i] = pb.NewNodeServiceClient(clientconn)
		}
	}
	return
}
func (node *Node) Service() {
	port := node.IpAddress[node.label-1]
	listener, err := net.Listen("tcp", port)
	if err != nil {
		node.Log.Fatalf("[Node %d] fail to listen:%v", node.label, err)
	}
	server := grpc.NewServer()
	pb.RegisterNodeServiceServer(server, node)
	reflection.Register(server)
	err = server.Serve(listener)
	if err != nil {
		node.Log.Fatalf("[Node %d] fail to provide service", node.label)
	}
	node.Log.Printf("[Node %d] now serve on %s    ", node.label, node.IpAddress[node.label-1])

	return
}
func NewForWeb(degree, label, counter int, metadatapath string, secretid int) (Node, error) {
	if label < 0 {
		return Node{}, errors.New("Label must be a non-negative number!")
	}
	//file, _ := os.Create(logPath + "/log" + strconv.Itoa(label))
	//defer file.Close()
	ipRaw := ReadIpList(metadatapath)[0 : counter+1]
	//fmt.Println(":????")
	bip := ipRaw[0]
	ipList := ipRaw[1 : counter+1]

	if counter < 0 {
		return Node{}, errors.New("Counter must be a non-negative number!")
	}
	randState := rand.New(rand.NewSource(time.Now().Local().UnixNano()))
	//fixedRandState := rand.New(rand.NewSource(int64(3)))
	dc := commitment.DLCommit{}
	dc.SetupFix()
	dpc := commitment.DLPolyCommit{}
	dpc.SetupFix(counter + 1)

	p := gmp.NewInt(0)
	p.SetString("57896044618658097711785492504343953926634992332820282019728792006155588075521", 10)
	lambda := make([]*gmp.Int, counter)
	// Calculate Lagrange Interpolation
	denominator := poly.NewConstant(1)
	tmp, _ := poly.NewPoly(1)
	tmp.SetCoeffWithInt(1, 1)
	//fmt.Println(time.Now())
	for i := 0; i < degree*2+1; i++ {
		tmp.GetPtrtoConstant().Neg(gmp.NewInt(int64(i + 1)))
		denominator.MulSelf(tmp)
	}
	for i := 0; i < degree*2+1; i++ {
		lambda[i] = gmp.NewInt(0)
		deno, _ := poly.NewPoly(0)
		tmp.GetPtrtoConstant().Neg(gmp.NewInt(int64(i + 1)))

		deno.Divide(denominator, tmp)
		deno.EvalMod(gmp.NewInt(0), p, lambda[i])
		inter := gmp.NewInt(0)
		deno.EvalMod(gmp.NewInt(int64(i+1)), p, inter)
		interInv := gmp.NewInt(0)
		interInv.ModInverse(inter, p)
		lambda[i].Mul(lambda[i], interInv)
		lambda[i].Mod(lambda[i], p)
	}
	//fmt.Println(time.Now())

	_0Shares := make([]*gmp.Int, counter)
	for i := 0; i < counter; i++ {
		_0Shares[i] = gmp.NewInt(0)
	}
	_0ShareCount := 0
	_0ShareSum := gmp.NewInt(0)

	recPoint := make([]*point.Point, counter)
	recCounter := 0

	secretShares := make([]*point.Point, counter)
	secretShares2 := make([]*point.Point, counter)
	//tmpPoly, err := poly.NewRand(degree, fixedRandState, p)

	//fmt.Println(time.Now())
	proPoly, _ := poly.NewPoly(degree)
	recPoly, _ := poly.NewPoly(degree)
	newPoly, _ := poly.NewPoly(degree)
	shareCnt := 0

	s0 := gmp.NewInt(0)

	oldPolyCmt := make([]*pbc.Element, counter)
	midPolyCmt := make([]*pbc.Element, counter)
	newPolyCmt := make([]*pbc.Element, counter)
	for i := 0; i < counter; i++ {
		oldPolyCmt[i] = dpc.NewG1()
		midPolyCmt[i] = dpc.NewG1()
		newPolyCmt[i] = dpc.NewG1()
	}

	zeroShareCmt := dc.NewG1()
	zeroPolyCmt := dpc.NewG1()
	zeroPolyWit := dpc.NewG1()

	zerosumShareCmt := make([]*pbc.Element, counter)
	zerosumPolyCmt := make([]*pbc.Element, counter)
	zerosumPolyWit := make([]*pbc.Element, counter)

	for i := 0; i < counter; i++ {
		zerosumShareCmt[i] = dc.NewG1()
		zerosumPolyCmt[i] = dpc.NewG1()
		zerosumPolyWit[i] = dpc.NewG1()
	}
	// amt
	amtPoly := make([]*poly.Poly, (counter+5)<<2)
	amtPolyCmt := make([]*pbc.Element, (counter+5)<<2)
	amtBasicPolyCmt := make([]*pbc.Element, (counter+5)<<2)
	amtStep := 0
	amtPolycmtSize := 0
	for (1 << amtPolycmtSize) <= counter+1 {
		amtPolycmtSize += 1
	}
	amtPolycmtSize += 1
	for i := 0; i < ((counter + 5) << 2); i++ {
		tmp2Poly, _ := poly.NewPoly(degree)
		amtPoly[i] = &tmp2Poly
		amtPolyCmt[i] = dpc.NewG1()
		amtBasicPolyCmt[i] = dpc.NewG1()
	}

	totMsgSize := 0
	s1 := time.Now()
	e1 := time.Now()
	s2 := time.Now()
	e2 := time.Now()
	s3 := time.Now()
	e3 := time.Now()

	amtflag1 := 0
	amtflag2 := 0

	clientConn := make([]*grpc.ClientConn, counter)
	nodeService := make([]pb.NodeServiceClient, counter)

	iniflag := true

	fileName := metadatapath + "/Screct" + strconv.Itoa(secretid) + "Node" + strconv.Itoa(label) + ".log"
	tmplogger, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModePerm)
	if err != nil {
		tmplogger, err = os.Create(fileName)
	}
	//os.Truncate(fileName, 0)
	log := log.New(tmplogger, "", log.LstdFlags)
	finish := false
	//log.
	//log.Printf("[Node %d] new done,ip = %s", label, ipList[label-1])
	return Node{
		secretid:        secretid,
		MetadataPath:    metadatapath,
		IpAddress:       ipList,
		ipOfBoard:       bip,
		degree:          degree,
		label:           label,
		counter:         counter,
		randState:       randState,
		dc:              &dc,
		dpc:             &dpc,
		p:               p,
		s0:              s0,
		lambda:          lambda,
		_0Shares:        _0Shares,
		_0ShareCount:    &_0ShareCount,
		_0ShareSum:      _0ShareSum,
		secretShares:    secretShares,
		secretShares2:   secretShares2,
		recPoint:        recPoint,
		recCounter:      &recCounter,
		recPoly:         &recPoly,
		proPoly:         &proPoly,
		newPoly:         &newPoly,
		shareCnt:        &shareCnt,
		oldPolyCmt:      oldPolyCmt,
		midPolyCmt:      midPolyCmt,
		newPolyCmt:      newPolyCmt,
		zeroShareCmt:    zeroShareCmt,
		zeroPolyCmt:     zeroPolyCmt,
		zeroPolyWit:     zeroPolyWit,
		zerosumPolyCmt:  zerosumPolyCmt,
		zerosumPolyWit:  zerosumPolyWit,
		zerosumShareCmt: zerosumShareCmt,
		totMsgSize:      &totMsgSize,
		s1:              &s1,
		e1:              &e1,
		s2:              &s2,
		e2:              &e2,
		s3:              &s3,
		e3:              &e3,
		clientConn:      clientConn,
		nodeService:     nodeService,
		iniflag:         &iniflag,
		Log:             log,
		amtStep:         amtStep,
		amtPolyCmt:      amtPolyCmt,
		amtPoly:         amtPoly,
		amtBasicPolyCmt: amtBasicPolyCmt,
		amtPolyCmtSize:  amtPolycmtSize,
		amtflag1:        amtflag1,
		amtflag2:        amtflag2,
		finish:          finish,
	}, nil

}
func (node *Node) DeleteServe() {
	node.Log.Printf("[Node %d] delete serve ", node.label)
	node.server.Stop()
	return
}
func (node *Node) ServeForWeb() {
	port := node.IpAddress[node.label-1]
	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("[Node %d] fail to listen:%v", node.label, err)
	}
	node.Log.Printf("[Node %d] now serve on %s   ", node.label, node.IpAddress[node.label-1])
	server := grpc.NewServer()
	pb.RegisterNodeServiceServer(server, node)
	reflection.Register(server)
	node.server = server
	err = server.Serve(listener)
	if err != nil {
		log.Fatalf("[Node %d] fail to provide service", node.label)
	}
	return
}
func (node *Node) Initsecret(ctx context.Context, msg *pb.InitMsg) (*pb.ResponseMsg, error) {
	node.Log.Printf("[Node %d] get initial polynomial", node.label)
	//fmt.Println("msg in the function ", msg)
	degree := int(msg.GetDegree())
	counter := int(msg.GetCounter())
	secretid := int(msg.GetSecretid())
	coeffbytes := msg.GetCoeff()
	//fmt.Println("coeffbytes", coeffbytes)
	//coeff := make([]*gmp.Int, degree+1)
	//for i := 0; i < degree+1; i++ {
	//	coeff[i] = gmp.NewInt(0)
	//	coeff[i].SetBytes(coeffbytes[i])
	//} //turn to char?
	node.store_secret(degree, counter, secretid, coeffbytes)
	return &pb.ResponseMsg{}, nil
}

func (node *Node) store_secret(degree int, counter int, secretid int, coeffbyte [][]byte) {
	db := common.GetDB()
	//向数据库中插入新纪录
	//for _, bytes := range coeffbyte {
	//	fmt.Println("yes bytes:", bytes)
	//}
	for i := 0; i < len(coeffbyte); i++ {
		data := coeffbyte[i]
		newSecretshare := model.Secretshare{
			SecretId: uint(secretid),
			UnitId:   uint(node.label),
			Degree:   degree,
			Counter:  counter,
			RowNum:   i,
			Data:     data,
		}
		//返回结果
		db.Create(&newSecretshare)
	}
	//dc := commitment.DLCommit{}
	//dc.SetupFix()
	//dpc := commitment.DLPolyCommit{}
	//dpc.SetupFix(counter)
	//p := gmp.NewInt(0)
	//p.SetString("57896044618658097711785492504343953926634992332820282019728792006155588075521", 10)

	//store to mysql
	return

}
func (node *Node) save_secret(degree int, counter int, secretid int, coeffbyte [][]byte) {
	db := common.GetDB()
	//向数据库中插入新纪录
	//for _, bytes := range coeffbyte {
	//	fmt.Println("yes bytes:", bytes)
	//}
	for i := 0; i < len(coeffbyte); i++ {
		var newsecretshare model.Secretshare
		db.Where("secret_id = ? and unit_id = ? and row_num =?", secretid, node.label, i).Find(&newsecretshare)
		newsecretshare.Data = coeffbyte[i]
		db.Save(newsecretshare)
	}
	return
	//store to mysql

}
func (node *Node) SetSecret() {
	//get secret from mysql
	//node.secretid = secretid
	secretid := node.secretid
	db := common.GetDB()
	var secretshares []model.Secretshare
	result := db.Where("secret_id =?", secretid).Where("unit_id", node.label).Find(&secretshares)
	rowNum := result.RowsAffected
	//var newsecretshare model.Secretshare
	//db.Where("secret_id = ? and unit_id = ?", secretid, node.label).First(&newsecretshare)
	//degree := newsecretshare.Degree
	//counter := newsecretshare.Counter
	//secretid := int(newsecretshare.SecretId)
	// 	assert(rowNum==degree*2+1)
	coeff := make([]*gmp.Int, rowNum)
	for i := 0; int64(i) < rowNum; i++ {
		var newsecretshare model.Secretshare
		db.Where("secret_id = ? and unit_id = ? and row_num =?", secretid, node.label, i).Find(&newsecretshare)
		//Data存放秘密份额,多项式
		Data := newsecretshare.Data

		coeff[i] = gmp.NewInt(0)
		coeff[i].SetBytes(Data)
	}

	node.Set(coeff)
	return
	//return newsecretshare
}

func (node *Node) Set(coeff []*gmp.Int) {
	tmpPoly, err := poly.NewPoly(len(coeff) - 1)
	tmpPoly.SetbyCoeff(coeff)
	if err != nil {
		panic("Error initializing random tmpPoly")
	}
	node.recPoly.ResetTo(tmpPoly)
	node.recPoly.ResetTo(tmpPoly)
	node.recPoly.EvalMod(gmp.NewInt(0), node.p, node.s0)
}

func (node *Node) AddSecret(ctx context.Context, msg *pb.RequestMsg) (*pb.ResponseMsg, error) {
	node.Log.Printf("[Node %d] add a empty data to SQL", node.label)
	db := common.GetDB()
	for i := 0; i < node.degree*2+1; i++ {
		newSecretshare := model.Secretshare{
			SecretId: uint(node.secretid),
			UnitId:   uint(node.label),
			Degree:   node.degree,
			Counter:  node.counter,
			RowNum:   i,
			Data:     gmp.NewInt(0).Bytes(),
		}
		//返回结果
		db.Create(&newSecretshare)
	}
	return &pb.ResponseMsg{}, nil
}
func (node *Node) DeleteSecret(ctx context.Context, msg *pb.RequestMsg) (*pb.ResponseMsg, error) {

	node.Log.Printf("[Node %d] delele secret's data in SQL", node.label)
	db := common.GetDB()
	for i := 0; i < node.degree*2+1; i++ {
		db.Where("secret_id = ? and unit_id = ? and row_num =?", node.secretid, node.label, i).Delete(&model.Secretshare{})
	}
	return &pb.ResponseMsg{}, nil
}
