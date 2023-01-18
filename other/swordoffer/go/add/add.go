package add

// 预备知识
// 有符号整数通常用补码来表示和存储，补码具有如下特征：
// - 正整数的补码与原码相同；负整数的补码为其原码除符号位外的所有位取反后加 1。
// - 可以将减法运算转化为补码的加法运算来实现。
// - 符号位与数值位可以一起参与运算。

func add(a int, b int) int {
	for b != 0 {
		carry := uint(a&b) << 1
		a ^= b
		b = int(carry)
	}
	return a
}
