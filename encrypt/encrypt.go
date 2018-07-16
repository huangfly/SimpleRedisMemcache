package encrypt

import (
	"log"
	"errors"
	"crypto/cipher"
	"github.com/qingche123/sm_crypto_golang/sm4"
)

var(
	
	 sm4Key  = []byte{0x40,0x40,0x40,0x40,0x40,0x40,0x40,0x40,0x51,0x51,0x51,0x51,0x51,0x51,0x51,0x51}//sm4对称加密算法的密钥
	 chiper cipher.Block
)

var strFormatError = errors.New("Source string format is error")

//初始化创建一个加密因子
func init(){
	chiper, _ = sm4.NewCipher(sm4Key)
}

//将16进制数据转成字符串显示
func hex2String(src []byte)string{
	l := len(src)
	var dst string
	hexStr := "0123456789ABCDEF";
	for i := 0; i < l; i++{
		dst = dst + string(hexStr[src[i] >> 4])
		dst = dst + string(hexStr[src[i] & 0x0f])
	}
	return dst
}

//将字符串转成16进制数据显示
func str2Hex(src string, srclen int)(*[]byte, error){
	dst := make([]byte, 0)
	for i:= 0; i < srclen; i++{
		var member byte
		tmp :=  src[i]
		if tmp >='0' && tmp <= '9'{
			member = (tmp - '0') << 4		
		} else if tmp>='a' && tmp<='f'{
			member = (tmp-'a'+0x0a)<<4	
		}  else if tmp>='A' && tmp<='F'{
			member = (tmp-'A'+0x0a)<<4
		}else{
			return &dst, strFormatError
		}
		i += 1
		tmp = src[i]
		if tmp >='0' && tmp <= '9'{
			member |= tmp - '0'	
		} else if tmp>='a' && tmp<='f'{
			member |= tmp-'a'+0x0a
		}  else if tmp>='A' && tmp<='F'{
			member |= tmp-'A'+0x0a
		}else{
			return &dst, strFormatError
		}
		dst = append(dst, member)
	}
	return &dst, nil
}

//给不足16的倍数的长度的数据填充0x00
func fillZero(src *[]byte){
	l := len(*src)
	zeroNum := 16- (l % 16)
	for i := 0; i < zeroNum; i++{
		*src = append(*src, 0x00)
	}
	log.Println("*src = ", *src)
}

//获取数据内容中0x00的个数
func getZeroLen(src []byte) int{
	num := 0
	//这个地方很奇怪range去遍历[]byte查找0x00会有问题计数num不准。
	for i := 0; i < len(src); i++{
		if src[i] == 0x00{
			num += 1
		}
	}
	return num;
}

//加密数据，将加密后的16进制数据用string返回
func Encrypt(src string)(string, error){
	source :=  []byte(src)
	fillZero(&source)
	endst := make([]byte, len(source))
	chiper.Encrypt(endst, source)
	dst := hex2String(endst)
	return dst, nil
}

//解密数据，将解密后的16进制数据用string返回
func Decode(src string)(string, error){
	ensrc, err := str2Hex(src, len(src))
	endst := make([]byte, len(*ensrc))
	if err != nil{
		return "", err
	}
	chiper.Decrypt(endst, *ensrc)
	dst := endst[:len(endst) - getZeroLen(endst)]
	return string(dst), nil
}