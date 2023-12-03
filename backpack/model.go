package backpack

import (
	"fmt"
	"sort"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
)

type Table struct {
	x_arr []int
	rows []Table_row
	
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

// For conditions
type Condition struct {
	max_weight int `json:"max_weight"`
	t_parts    []T_part `json:"t"`
	columns_num  int `json: "columns_num"`
	tables []Table
}

type T_part struct {
	m     int
	r     int
	index int
}

func (c *Condition) CreateCondition() {

	// Insering number of T columns
	fmt.Println("Enter number of T")
	n := 0
	fmt.Scan(&n)
	c.columns_num = n

	// Inserting m and r into T tables
	c.t_parts = make([]T_part, n)
	for i := 1; i <= n; i++ {
		var m, r int
		fmt.Printf("Enter m and r for T%d\n", i)
		fmt.Scan(&m, &r)
		c.t_parts[i - 1] = T_part{m, r, i}
	}
	
	fmt.Println("Enter max weight: ")
	fmt.Scan(&c.max_weight)
}


// Condition functions
func (c *Condition) Solve() int{
	max := 0
	c.tables = make([]Table, c.columns_num)

	x_next := []int{ 0 }

	for i := 0; i < c.columns_num; i++{
		c.tables[i] = Table{
			x_arr: x_next,
		}
		x_next = c.tables[i].SolveX(c.max_weight, c.t_parts[i])
		x_next = removeDuplicate[int](x_next)
		sort.Slice(x_next, func(i, j int) bool {
			return x_next[i] < x_next[j]
		})
	}

	for i := len(c.tables) - 1; i >= 0; i--{
		for j := 0; j < len(c.tables[i].rows); j++{
			c.tables[i].rows[j].SolveB()
		}
		// for _, r := range c.tables[i].rows{
		// 	r.SolveB()
		// 	fmt.Println(r.B_n)
		// }
	}

	return max
}

func (c * Condition) PrintTables(){
	for i, tabl := range c.tables{
		// fmt.Println("==================================================================")
		// fmt.Printf("                                 T%d\n", i + 1)
		// fmt.Println("==================================================================")

		tab := table.NewWriter()
		tab.SetTitle(fmt.Sprintf("T%d\n", i + 1))
		tab.Style().Format.Header = text.FormatTitle
		tab.AppendHeader(table.Row{fmt.Sprintf("x%d", i), fmt.Sprintf("u%d", i + 1), fmt.Sprintf("x%d", i + 1), fmt.Sprintf("z%d", i + 1), fmt.Sprintf("B%d(x%d)", i + 1, i + 1), fmt.Sprintf("z%d + B%d", i + 1, i + 1), fmt.Sprintf("B%d(x%d)", i, i)})
		
		for _, r := range tabl.rows{
			for k := 0; k < len(r.u_n); k++{
				if k == 0{
					//tab.AppendRow(table.Row{r.x_past, r.u_n[k], r.x_n[k], r.z_n[k], r.B_n[k], r.B_z[k], r.B_past})
					tab.AppendRow(table.Row{r.x_past, r.u_n[k], r.x_n[k], r.z_n[k], r.B_n[k], r.B_z[k], r.B_past})
					continue
				}
				tab.AppendRow(table.Row{" ", r.u_n[k], r.x_n[k], r.z_n[k], r.B_n[k], r.B_z[k], " "})

			}

		}
		fmt.Println(tab.Render())
	}
}


// Table functions
func (t *Table) SolveX(max_weight int, t_cond T_part) []int{

	ans := []int{}

	t.rows = make([]Table_row, len(t.x_arr))
	
	for i, tab := range t.x_arr{
		t.rows[i].x_past = tab
		t.rows[i].SolveByMax(max_weight, t_cond.m, t_cond.r)
		ans = append(ans, t.rows[i].x_n...)
	}

	return ans
}



// Table row functions
func (t *Table_row) SolveByMax(max, weight_step, price_step int) {

	// First row of every string is zeros
	u_temp := 0
	x_temp := t.x_past
	z_temp := 0
	t.x_n = append(t.x_n, x_temp)
	t.u_n = append(t.u_n, 0)
	t.z_n = append(t.z_n, 0)

	// Running while row's x lower then max
	for x_temp + weight_step <= max{
		u_temp++
		x_temp += weight_step
		z_temp += price_step
		t.x_n = append(t.x_n, x_temp)
		t.u_n = append(t.u_n, u_temp)
		t.z_n = append(t.z_n, z_temp)
	}
}


var B_x map[int]int = make(map[int]int)


func (t * Table_row) SolveB(){
	b_n_temp := 0
	b_z_temp := 0
	max_bz := 0
	for i := 0; i < len(t.x_n); i++{
		if a, found := B_x[t.x_n[i]]; found{
			b_n_temp = a
			b_z_temp = b_n_temp + t.z_n[i]
			t.B_n = append(t.B_n, b_n_temp)
			t.B_z = append(t.B_z, b_z_temp)
			if b_z_temp >= max_bz {
				max_bz = b_z_temp
			}

			delete(B_x, t.x_n[i])
			
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
