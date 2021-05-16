package mysql

import (
	"fmt"
	"testing"
	"time"
)

// TestNewConfig
func TestNewConfig(t *testing.T) {
	var config = NewConfig("", "", "", "")
	if _, ok := interface{}(config).(Config); !ok {
		t.Errorf("instance error")
		return
	}
}

// TestConfig_ToDSN
func TestConfig_ToDSN(t *testing.T) {
	var dstStr = "admin:mEurV@3457cA@tcp(3.4.5.6:3306)/?loc=Asia%2FShanghai&charset=UTF8MB4"
	var config = NewConfig("3.4.5.6:3306", "admin", "mEurV@3457cA", "")
	if config.ToDSN() != dstStr {
		t.Errorf("ToDSN error")
		return
	}
}

// TestWithLoc
func TestWithLoc(t *testing.T) {
	var dstStr = "admin:123456@tcp(3.4.5.6:3306)/?loc=Asia%2FChongQing&charset=UTF8MB4"
	var config = NewConfig("3.4.5.6:3306", "admin", "123456", "",
		WithLoc("Asia/ChongQing"))
	if config.ToDSN() != dstStr {
		t.Errorf("WithLoc error")
		return
	}
}

// TestWithCharset
func TestWithCharset(t *testing.T) {
	var dstStr = "admin:123456@tcp(3.4.5.6:3306)/?loc=Asia%2FShanghai&charset=UTF8"
	var config = NewConfig("3.4.5.6:3306", "admin", "123456", "",
		WithCharset("UTF8"))
	if config.ToDSN() != dstStr {
		t.Errorf("WithCharset error")
		return
	}
}

// TestWithProtocol
func TestWithProtocol(t *testing.T) {
	var dstStr = "admin:123456@unix(3.4.5.6:3306)/?loc=Asia%2FShanghai&charset=UTF8MB4"
	var config = NewConfig("3.4.5.6:3306", "admin", "123456", "",
		WithProtocol("unix"))
	if config.ToDSN() != dstStr {
		t.Errorf("WithProtocol error")
		return
	}
}

// TestWithAllowAllFiles
func TestWithAllowAllFiles(t *testing.T) {
	var dstStr = "admin:123456@tcp(3.4.5.6:3306)/?loc=Asia%2FShanghai&charset=UTF8MB4&allowAllFiles=true"
	var config = NewConfig("3.4.5.6:3306", "admin", "123456", "",
		WithAllowAllFiles())
	if config.ToDSN() != dstStr {
		t.Errorf("WithAllowAllFiles error")
		return
	}
}

// TestWithAllowCleartextPasswords
func TestWithAllowCleartextPasswords(t *testing.T) {
	var dstStr = "admin:123456@tcp(3.4.5.6:3306)/?loc=Asia%2FShanghai&charset=UTF8MB4&allowCleartextPasswords=true"
	var config = NewConfig("3.4.5.6:3306", "admin", "123456", "",
		WithAllowCleartextPasswords())
	if config.ToDSN() != dstStr {
		t.Errorf("WithAllowCleartextPasswords error")
		return
	}
}

// TestWithAllowNativePasswords
func TestWithAllowNativePasswords(t *testing.T) {
	var dstStr = "admin:123456@tcp(3.4.5.6:3306)/?loc=Asia%2FShanghai&charset=UTF8MB4&allowNativePasswords=true"
	var config = NewConfig("3.4.5.6:3306", "admin", "123456", "",
		WithAllowNativePasswords())
	if config.ToDSN() != dstStr {
		t.Errorf("WithAllowNativePasswords error")
		return
	}
}

// WithAllowOldPasswords
func TestWithAllowOldPasswords(t *testing.T) {
	var dstStr = "admin:123456@tcp(3.4.5.6:3306)/?loc=Asia%2FShanghai&charset=UTF8MB4&allowOldPasswords=true"
	var config = NewConfig("3.4.5.6:3306", "admin", "123456", "",
		WithAllowOldPasswords())
	if config.ToDSN() != dstStr {
		t.Errorf("WithAllowOldPasswords error")
		return
	}
}

// TestWithCheckConnLiveNess
func TestWithCheckConnLiveNess(t *testing.T) {
	var dstStr = "admin:123456@tcp(3.4.5.6:3306)/?loc=Asia%2FShanghai&charset=UTF8MB4&checkConnLiveness=true"
	var config = NewConfig("3.4.5.6:3306", "admin", "123456", "",
		WithCheckConnLiveNess())
	if config.ToDSN() != dstStr {
		t.Errorf("WithCheckConnLiveNess error")
		return
	}
}

// TestWithCollation
func TestWithCollation(t *testing.T) {
	var dstStr = "admin:123456@tcp(3.4.5.6:3306)/?loc=Asia%2FShanghai&charset=UTF8MB4&collation=123"
	var config = NewConfig("3.4.5.6:3306", "admin", "123456", "",
		WithCollation("123"))
	if config.ToDSN() != dstStr {
		t.Errorf("WithCollation error")
		return
	}
}

// TestWithClientFoundRows
func TestWithClientFoundRows(t *testing.T) {
	var dstStr = "admin:123456@tcp(3.4.5.6:3306)/?loc=Asia%2FShanghai&charset=UTF8MB4&clientFoundRows=true"
	var config = NewConfig("3.4.5.6:3306", "admin", "123456", "",
		WithClientFoundRows())
	if config.ToDSN() != dstStr {
		t.Errorf("WithClientFoundRows error")
		return
	}
}

// TestWithColumnsWithAlias
func TestWithColumnsWithAlias(t *testing.T) {
	var dstStr = "admin:123456@tcp(3.4.5.6:3306)/?loc=Asia%2FShanghai&charset=UTF8MB4&columnsWithAlias=true"
	var config = NewConfig("3.4.5.6:3306", "admin", "123456", "",
		WithColumnsWithAlias())
	if config.ToDSN() != dstStr {
		t.Errorf("WithColumnsWithAlias error")
		return
	}
}

// TestWithInterpolateParams
func TestWithInterpolateParams(t *testing.T) {
	var dstStr = "admin:123456@tcp(3.4.5.6:3306)/?loc=Asia%2FShanghai&charset=UTF8MB4&interpolateParams=true"
	var config = NewConfig("3.4.5.6:3306", "admin", "123456", "",
		WithInterpolateParams())
	if config.ToDSN() != dstStr {
		t.Errorf("TestWithInterpolateParams error")
		return
	}
}

// TestWithInterpolateParams
func TestWithMaxAllowedPacket(t *testing.T) {
	var dstStr = "admin:123456@tcp(3.4.5.6:3306)/?loc=Asia%2FShanghai&charset=UTF8MB4&maxAllowedPacket=65535"
	var config = NewConfig("3.4.5.6:3306", "admin", "123456", "",
		WithMaxAllowedPacket(65535))
	if config.ToDSN() != dstStr {
		t.Errorf("WithMaxAllowedPacket error")
		return
	}
}

// TestWithMultiStatements
func TestWithMultiStatements(t *testing.T) {
	var dstStr = "admin:123456@tcp(3.4.5.6:3306)/?loc=Asia%2FShanghai&charset=UTF8MB4&multiStatements=true"
	var config = NewConfig("3.4.5.6:3306", "admin", "123456", "",
		WithMultiStatements())
	if config.ToDSN() != dstStr {
		t.Errorf("WithMaxAllowedPacket error")
		return
	}
}

// TestWithParseTime
func TestWithParseTime(t *testing.T) {
	var dstStr = "admin:123456@tcp(3.4.5.6:3306)/?loc=Asia%2FShanghai&charset=UTF8MB4&parseTime=true"
	var config = NewConfig("3.4.5.6:3306", "admin", "123456", "",
		WithParseTime())
	if config.ToDSN() != dstStr {
		t.Errorf("WithParseTime error")
		return
	}
}

// TestWithReadTimeout
func TestWithReadTimeout(t *testing.T) {
	var dstStr = fmt.Sprintf("admin:123456@tcp(3.4.5.6:3306)/?loc=Asia%%2FShanghai&charset=UTF8MB4&readTimeout=%d",
		time.Duration(112233).Milliseconds())

	var config = NewConfig("3.4.5.6:3306", "admin", "123456", "",
		WithReadTimeout(112233))
	if config.ToDSN() != dstStr {
		t.Errorf("WithReadTimeout error")
		return
	}
}

// TestWithRejectReadOnly
func TestWithRejectReadOnly(t *testing.T) {
	var dstStr = "admin:123456@tcp(3.4.5.6:3306)/?loc=Asia%2FShanghai&charset=UTF8MB4&rejectReadOnly=true"
	var config = NewConfig("3.4.5.6:3306", "admin", "123456", "",
		WithRejectReadOnly())
	if config.ToDSN() != dstStr {
		t.Errorf("WithRejectReadOnly error")
		return
	}
}

// TestWithServerPubKey
func TestWithServerPubKey(t *testing.T) {
	var dstStr = "admin:123456@tcp(3.4.5.6:3306)/?loc=Asia%2FShanghai&charset=UTF8MB4&serverPubKey=Volibear"
	var config = NewConfig("3.4.5.6:3306", "admin", "123456", "",
		WithServerPubKey("Volibear"))
	if config.ToDSN() != dstStr {
		t.Errorf("WithServerPubKey error")
		return
	}
}

// TestWithTimeout
func TestWithTimeout(t *testing.T) {
	var dstStr = fmt.Sprintf("admin:123456@tcp(3.4.5.6:3306)/?loc=Asia%%2FShanghai&charset=UTF8MB4&timeout=%d",
		time.Duration(65535).Milliseconds())

	var config = NewConfig("3.4.5.6:3306", "admin", "123456", "",
		WithTimeout(65535))
	if config.ToDSN() != dstStr {
		t.Errorf("WithTimeout error")
		return
	}
}

// TestWithWriteTimeout
func TestWithWriteTimeout(t *testing.T) {
	var dstStr = fmt.Sprintf("admin:123456@tcp(3.4.5.6:3306)/?loc=Asia%%2FShanghai&charset=UTF8MB4&writeTimeout=%d",
		time.Duration(65535).Milliseconds())

	var config = NewConfig("3.4.5.6:3306", "admin", "123456", "",
		WithWriteTimeout(65535))
	if config.ToDSN() != dstStr {
		t.Errorf("WithWriteTimeout error")
		return
	}
}

// TestWithMaxOpenConn
func TestWithMaxOpenConn(t *testing.T) {

	var config = NewConfig("3.4.5.6:3306", "admin", "123456", "",
		WithMaxOpenConn(65535))

	if config.MaxOpenConn != 65535 {
		t.Errorf("WithMaxOpenConn error")
		return
	}
}

// TestWithMaxIdleConn
func TestWithMaxIdleConn(t *testing.T) {

	var config = NewConfig("3.4.5.6:3306", "admin", "123456", "",
		WithMaxIdleConn(65535))

	if config.MaxIdleConn != 65535 {
		t.Errorf("WithMaxIdleConn error")
		return
	}
}

// TestWithConnMaxLifeTime
func TestWithConnMaxLifeTime(t *testing.T) {

	var config = NewConfig("3.4.5.6:3306", "admin", "123456", "",
		WithConnMaxLifeTime(time.Second))

	if config.ConnMaxLifeTime != time.Second {
		t.Errorf("WithConnMaxLifeTime error")
		return
	}
}

// TestWithConnMaxIdleTime
func TestWithConnMaxIdleTime(t *testing.T) {

	var config = NewConfig("3.4.5.6:3306", "admin", "123456", "",
		WithConnMaxIdleTime(time.Second))

	if config.ConnMaxIdleTime != time.Second {
		t.Errorf("WithConnMaxIdleTime error")
		return
	}
}

// TestWithTls
func TestWithTls(t *testing.T) {
	var dstStr = "admin:123456@tcp(3.4.5.6:3306)/?loc=Asia%2FShanghai&charset=UTF8MB4&tls=preferred"
	var config = NewConfig("3.4.5.6:3306", "admin", "123456", "",
		WithTls("preferred"))

	if config.ToDSN() != dstStr {
		t.Errorf("WithTls error")
		return
	}
}
