# SimpleRedisMemcache
用golang实现的一个可以远程的数据库缓存系统，提供了key-value的数据缓存功能以及数据的增删改查的基本操作，配置文件通过命令行可选。
同时为了数据安全性以及配置文件中包含的数据库和reids密码的安全性采用了sm4对称加密算法加密的功能。
mian包中实现了一个基本的demo通过json进行缓存的数据传输。
缓存的数据结构提供了最基本的key-value有需要可以根据业务进行数据结构的修改或者提供其他的业务操作，基本的框架不变，希望多交流，多探讨，多指正。