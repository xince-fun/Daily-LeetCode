# 树上倍增算法


[1483. 树节点的第k个祖先](https://leetcode.cn/problems/kth-ancestor-of-a-tree-node/solutions/2305895/mo-ban-jiang-jie-shu-shang-bei-zeng-suan-v3rw/)

预处理出每个节点的第 $2^i$ 个祖先节点，即第 $2,4,8, \cdots$ 个祖先节点（注意 $x$ 的第 1 个祖先节点就是 $parent[x]$）。由于任意 $k$ 可以分解为若干不同的 $2$ 的幂（例如 $13=8+4+1$），所以只需要预处理出这些 $2^i$ 祖先节点，就可以快速地到达任意第 $k$ 个祖先节点。

例如 $k=13=8+4+1=1101_{(2)}$，可以先往上跳 $8$ 步，再往上跳 $4$ 步和 $1$ 步；也可以先往上跳 $1$ 步，再往上跳 $4$ 步和 $8$ 步。无论如何跳，都只需要跳 $3$ 次就能到达第 $13$ 个祖先节点。

据此，可以得到下面的算法。

### 算法

在构造函数 $TreeAncestor$中，预处理出每个节点 $x$ 的第 $2^i$ 个祖先节点，计作 $pa[x][i]$ （若第 $2^i$个祖先节点不存在，则 $pa[x][i] = -1$）

- 先枚举 $i$, 再枚举$x$。相当于先算出所有爷爷节点，再算出所有爷爷的爷爷节点，依次类推。
- $pa[x][0] = parent[x]$，即父节点。
- $pa[x][1] = pa[pa[x][0]][0]$， 即爷爷节点。
- 依次类推， $pa[x][i+1] = pa[pa[x][i]][i]$，表示$x$的第$2^i$个祖先节点的第$2^i$个祖先节点就是 $x$的第$2^{i+1}$个祖先节点。特别地，如果 $pa[x][i] = -1$ 则 $pa[x][i+1] = -1$。
- 这里 $i+1$至多为  $ \lfloor \log_2 n \rfloor$。例如 $n=13$时，$ \lfloor \log_2 13 \rfloor = 3$，至多需要预处理到第 $2^3$个祖先节点。 （当然，你也可以先把树高，或者每个节点的深度求出来，再据此做精细地计算。）

对于 $getKthAncestor$，需要找到 $k$ 的二进制表示中的所有 $1$（相当于把 $k$ 分解为若干 $2^i$ ）。可以从小到大枚举 $i$，如果 $k$ 右移 $i$ 位后的最低位为 $1$，就说明 $k$ 的二进制从低到高第 $i$ 位是 $1$，那么往上跳 $2^i$ 步，将 $node$更新为 $pa[node][i]$。如果 $node=−1$ 则说明第 $k$ 个祖先节点不存在。


```go
type TreeAncestor [][]int

func Constructor(n int, parent []int) TreeAncestor {
    m := bits.Len(uint(n))
    pa := make([][]int, n)
    for i, p := range parent {
        pa[i] = make([]int, m)
        pa[i][0] = p
    }
    for i := 0; i < m-1; i++ {
        for x := 0; x < n; x++ {
            if p := pa[x][i]; p != -1 {
                pa[x][i+1] = pa[p][i]
            } else {
                pa[x][i+1] = -1
            }
        }
    }
    return pa
}

func (pa TreeAncestor) GetKthAncestor(node, k int) int {
    m := bits.Len(uint(k))
    for i := 0; i < m; i++ {
        if k>>i&1 > 0 { // k 的二进制从低到高第 i 位是 1
            node = pa[node][i]
            if node < 0 {
                break
            }
        }
    }
    return node
}

// 另一种写法，不断去掉 k 的最低位的 1
func (pa TreeAncestor) GetKthAncestor2(node, k int) int {
    for ; k > 0 && node != -1; k &= k - 1 {
        node = pa[node][bits.TrailingZeros(uint(k))]
    }
    return node
}
```