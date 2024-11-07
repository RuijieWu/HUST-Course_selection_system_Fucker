/*
 * @Author: 7erry
 * @Date: 2024-11-07 13:19:33
 * @LastEditTime: 2024-11-07 13:41:33
 * @Description:
 */
package client

import (
	"fmt"
	"testing"
)

func TestGetTimeDiff(t *testing.T) {
	f := NewFucker()
	fmt.Println(f.GetTimeDiff())
}
