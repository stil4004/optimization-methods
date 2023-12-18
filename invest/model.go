package invest

import (
	"fmt"
	"os"
	"sort"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
)

type Condition struct {
	deposit, num_of_companies int
	Companies                 []Company
	tables                    []Table
}

type Company struct{
	prices []int
}

type Table struct {
	x_arr []int
	rows  []Table_row
	
}

type Table_row struct{
	x_past int
	u_n []int
	x_n []int
	z_n []int
	B_n []int
	B_z []int
	B_past int
}

func (c *Condition) Input() {

	fmt.Println("Enter deposite: ")
	fmt.Scan(&c.deposit)

	fmt.Println("Enter number of companies")
	fmt.Scan(&c.num_of_companies)

	for j := 0; j < c.num_of_companies; j++{
		c.Companies = append(c.Companies, Company{
			prices: []int{},
		})
	}

	for i := 0; i <= c.deposit; i++{
		//temp_arr := []int{}
		for j := 0; j < c.num_of_companies; j++{
			fmt.Printf("Input profit for %d company to %d deposit: \n", j + 1, i)
			var new_price int
			fmt.Scan(&new_price)
			c.Companies[j].prices = append(c.Companies[j].prices, new_price)	
		}
	}
	fmt.Println(c)
}

func (c * Condition) Solve() int{
	max := 0
	c.tables = make([]Table, c.num_of_companies)

	x_next := []int{ 0 }
	isLast := false

	for i := 0; i < c.num_of_companies; i++{
		c.tables[i] = Table{
			x_arr: x_next,
		}

		if i == c.num_of_companies - 1{
			isLast = true
		}
		x_next = c.tables[i].SolveX(c.deposit, c.Companies[i], isLast)
		x_next = removeDuplicate[int](x_next)
		sort.Slice(x_next, func(i, j int) bool {
			return x_next[i] < x_next[j]
		})
	}

	for i := len(c.tables) - 1; i >= 0; i--{
		for j := 0; j < len(c.tables[i].rows); j++{
			c.tables[i].rows[j].SolveB()
		}
	}

	return max
}

func (c * Condition) PrintTables(){
	//os.RemoveAll("../bin/answers/")
	f, err := os.OpenFile("bin/answers/invest_answer.txt", os.O_RDWR | os.O_CREATE | os.O_TRUNC, os.ModePerm)
	if err != nil{
		panic(err)
	}
	defer f.Close()
	
	for i, tabl := range c.tables{
		tab := table.NewWriter()
		tab.SetTitle(fmt.Sprintf("T%d\n", i + 1))
		tab.Style().Format.Header = text.FormatTitle
		tab.AppendHeader(table.Row{fmt.Sprintf("x%d", i), fmt.Sprintf("u%d", i + 1), fmt.Sprintf("x%d", i + 1), fmt.Sprintf("z%d", i + 1), fmt.Sprintf("B%d(x%d)", i + 1, i + 1), fmt.Sprintf("z%d + B%d", i + 1, i + 1), fmt.Sprintf("B%d(x%d)", i, i)})
		
		for _, r := range tabl.rows{
			for k := 0; k < len(r.u_n); k++{
				if k == 0{
					tab.AppendRow(table.Row{r.x_past, r.u_n[k], r.x_n[k], r.z_n[k], r.B_n[k], r.B_z[k], r.B_past})
					continue
				}
				tab.AppendRow(table.Row{" ", r.u_n[k], r.x_n[k], r.z_n[k], r.B_n[k], r.B_z[k], " "})

			}

		}
		fmt.Println(tab.Render())
		fmt.Fprintf(f, "%s\n", tab.Render())
	}
}

func (c *Condition) FPrintTables(){
	f, err := os.OpenFile("../bin/answer.txt", os.O_CREATE | os.O_RDWR, os.ModeAppend)
	if err != nil{
		panic(err)
	}
	
	for i, tabl := range c.tables{
		tab := table.NewWriter()
		tab.SetTitle(fmt.Sprintf("T%d\n", i + 1))
		tab.Style().Format.Header = text.FormatTitle
		tab.AppendHeader(table.Row{fmt.Sprintf("x%d", i), fmt.Sprintf("u%d", i + 1), fmt.Sprintf("x%d", i + 1), fmt.Sprintf("z%d", i + 1), fmt.Sprintf("B%d(x%d)", i + 1, i + 1), fmt.Sprintf("z%d + B%d", i + 1, i + 1), fmt.Sprintf("B%d(x%d)", i, i)})
		
		for _, r := range tabl.rows{
			for k := 0; k < len(r.u_n); k++{
				if k == 0{
					tab.AppendRow(table.Row{r.x_past, r.u_n[k], r.x_n[k], r.z_n[k], r.B_n[k], r.B_z[k], r.B_past})
					continue
				}
				tab.AppendRow(table.Row{" ", r.u_n[k], r.x_n[k], r.z_n[k], r.B_n[k], r.B_z[k], " "})

			}

		}
		fmt.Println(tab.Render())
		fmt.Fprintf(f, "%s\n", tab.Render())
	}
}

func (t *Table) SolveX(max int, comp Company, isLast bool) []int {

	ans := []int{}

	t.rows = make([]Table_row, len(t.x_arr))
	
	for i, tab := range t.x_arr{
		t.rows[i].x_past = tab
		t.rows[i].SolveByMax(max, comp.prices, isLast)
		ans = append(ans, t.rows[i].x_n...)
	}

	return ans

}

// Table row functions
func (t *Table_row) SolveByMax(max int, prices []int, isLast bool) {
	if isLast{
		u_temp := max - t.x_past
		x_temp := max
		z_temp := 0
		//if u_temp == 0 {
		//	z_temp = 0
		//} else {
		z_temp = prices[u_temp]
		//}
		t.x_n = append(t.x_n, x_temp)
		t.u_n = append(t.u_n, u_temp)
		t.z_n = append(t.z_n, z_temp)
		return
	}

	// First row of every string is zeros
	i := 0
	u_temp := 0
	x_temp := t.x_past
	z_temp := prices[u_temp]
	t.x_n = append(t.x_n, x_temp)
	t.u_n = append(t.u_n, 0)
	t.z_n = append(t.z_n, z_temp)

	// Running while row's x lower then max
	for x_temp + 1 <= max{
		u_temp++
		x_temp += 1
		z_temp = prices[u_temp]
		t.x_n = append(t.x_n, x_temp)
		t.u_n = append(t.u_n, u_temp)
		t.z_n = append(t.z_n, z_temp)
		i++
	}
}

var B_x map[int]int = make(map[int]int)

func (t * Table_row) SolveB(){
	b_n_temp := 0
	b_z_temp := 0
	max_bz := -10000
	for i := 0; i < len(t.x_n); i++{
		if a, found := B_x[t.x_n[i]]; found{
			b_n_temp = a
			b_z_temp = b_n_temp + t.z_n[i]
			t.B_n = append(t.B_n, b_n_temp)
			t.B_z = append(t.B_z, b_z_temp)
			if b_z_temp >= max_bz {
				max_bz = b_z_temp
			}

			//delete(B_x, t.x_n[i])
			
			continue
		}
		b_n_temp = 0
		b_z_temp = b_n_temp + t.z_n[i]
		t.B_n = append(t.B_n, b_n_temp)
		t.B_z = append(t.B_z, b_z_temp)
		if b_z_temp > max_bz {
			max_bz = b_z_temp
		}

	}
	t.B_past = max_bz
	B_x[t.x_past] = max_bz

}

// Helpfull functions
func removeDuplicate[T string | int](sliceList []T) []T {
    allKeys := make(map[T]bool)
    list := []T{}
    for _, item := range sliceList {
        if _, value := allKeys[item]; !value {
            allKeys[item] = true
            list = append(list, item)
        }
    }
    return list
}