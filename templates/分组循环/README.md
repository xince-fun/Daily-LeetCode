# 分组循环

一般来说，分组循环的模板如下

```python
n = len(nums)
i = 0
while i < n:
    start = i
    while i < n and ...:
        i += 1
    # 从 start 到 i-1 是一组
    # 下一组从 i 开始， 无需 i += 1
```