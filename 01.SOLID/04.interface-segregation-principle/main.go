package main

type Document struct {
}

type Machine interface {
	Print(d *Document)
	Fax(d *Document)
	Scan(d *Document)
}

type MultiFunctionPrinter struct {
}

func (m *MultiFunctionPrinter) Print(d *Document) {
}

func (m *MultiFunctionPrinter) Fax(d *Document) {
}

func (m *MultiFunctionPrinter) Scan(d *Document) {
}

type OldFashionedPrinter struct {
}

func (o *OldFashionedPrinter) Print(d *Document) {
}

func (o *OldFashionedPrinter) Fax(d *Document) {
	panic("operation not supported")
}

// Deprecated: ...
func (o *OldFashionedPrinter) Scan(d *Document) {
	panic("operation not supported")
}

// ISP
type Printer interface {
	Print(d *Document)
}

type Scanner interface {
	Scan(d *Document)
}

type MyPrinter struct {
}

func (m *MyPrinter) Print(d *Document) {
}

type Photocopier struct {
}

func (p *Photocopier) Print(d *Document) {
}

func (p *Photocopier) Scan(d *Document) {
}

type MultiFunctionDevice interface {
	Printer
	Scanner
}

func main() {
	ofp := OldFashionedPrinter{}
	ofp.Print(&Document{})
	ofp.Scan(&Document{}) // IDE will warn you that this method is deprecated
}

// ISP - Interface Segregation Principle
// Interface segregation principle states that no client should be forced to depend on methods it does not use.

// In the example above, we have a Machine interface that has three methods: Print, Fax and Scan. This interface is implemented by MultiFunctionPrinter and OldFashionedPrinter.
// MultiFunctionPrinter implements all three methods, but OldFashionedPrinter implements only Print and Fax. The problem is that OldFashionedPrinter is forced to implement Scan method, even though it does not need it.
// The solution is to split the Machine interface into smaller interfaces. In this case, we can split it into three interfaces: Printer, Scanner and Faxer. Then, we can implement these interfaces in our types.
// The main advantage of this approach is that we can use interfaces as types. This means that we can pass Printer, Scanner or Faxer to a function, instead of passing the whole Machine interface.
// This is useful because we can pass only the methods that we need. For example, we can pass only the Printer interface to a function that only needs to print documents.
