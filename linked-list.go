package main
import ( 
	"fmt"
	u "github.com/italorfeitosa/data-structures-with-go/utils"
)

type Block struct {
	ID string
	prev *Block
}

type BlockChain struct {
	lastBlockAdded *Block
	length int
}

func newBlockChain() *BlockChain {
	return &BlockChain{
		lastBlockAdded: nil,
		length: 0,
	}
}

func (bc *BlockChain) addBlock(b *Block) {
	if bc.lastBlockAdded != nil {
		b.prev = bc.lastBlockAdded
		bc.lastBlockAdded = b
	} else {
		bc.lastBlockAdded = b
	}
	bc.length += 1
}

func newBlock() *Block{
	id := u.GenerateUuid();
	u.Log(id)
	return &Block{
		ID: id,
		prev: nil,
	}
}

func main() {
	blockChain := newBlockChain()
	blockChain.addBlock(newBlock())
	blockChain.addBlock(newBlock())
	blockChain.addBlock(newBlock())
	blockChain.addBlock(newBlock())
	blockChain.addBlock(newBlock())
	fmt.Println(blockChain.lastBlockAdded.prev)
}