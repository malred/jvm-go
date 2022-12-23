package classpath

import (
	"os"
	"strings"
)
import "path/filepath"

// WildcardEntry实际上也是CompositeEntry(命令行使用了 * )
func newWildcardEntry(path string) CompositeEntry {
	// 把路径末尾的星号去掉，得到baseDir
	baseDir := path[:len(path)-1] // remove *
	compositeEntry := []Entry{}
	//根据后缀名选出JAR文件，并且返回SkipDir跳过子目录（通配符类路径不能递归匹配子目录下的JAR文件）
	walkFn := func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		// 如果info(os.FileInfo)是目录,并且path不是baseDir
		if info.IsDir() && path != baseDir {
			// 跳过该目录
			return filepath.SkipDir
		}
		// 如果后缀是.jar或.JAR
		if strings.HasSuffix(path, ".jar") || strings.HasSuffix(path, ".JAR") {
			jarEntry := newZipEntry(path)                     // 让zipEntry解析
			compositeEntry = append(compositeEntry, jarEntry) // 添加到compositionEntry
		}
		return nil
	}
	// 调用filepath包的Walk函数遍历baseDir创建ZipEntry
	filepath.Walk(baseDir, walkFn)
	return compositeEntry
}
