package main
import ( 
	u "github.com/italorfeitosa/data-structures-with-go/utils"
)

type Block struct {
	ID string
	content interface{}
	next *Block
}

type BlockChain struct {
	head *Block
	tail *Block
	length int
}

func newBlockChain() *BlockChain {
	return &BlockChain{
		head: nil,
		length: 0,
	}
}

func (bc *BlockChain) addBlock(b *Block) {
	if bc.head != nil {
		b.next = bc.head
		bc.head = b
	} else {
		bc.head = b
	}
	bc.length += 1
}

func newBlock(content interface{}) *Block{
	id := u.GenerateUuid();
	u.Log(id)
	return &Block{
		ID: id,
		content: content,
		next: nil,
	}
}

func main() {
	blockChain := newBlockChain()
	blockChain.addBlock(newBlock(nil))
	blockChain.addBlock(newBlock(42))
	blockChain.addBlock(newBlock("teste"))
	blockChain.addBlock(newBlock(32))
	blockChain.addBlock(newBlock(5+3))
	u.Log(blockChain.head)
}