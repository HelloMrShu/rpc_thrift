namespace go user
namespace php user

struct UserInfo {
    #序号:字段类型 字段名
    1:i64 id
    2:string name
}

# 定义一个用户服务
service UserService{
    # 定义一个GetUser方法（接收一个用户id，返回上面定义的用户信息）
    UserInfo GetUser(1:i32 id)
    # 定义一个GetName方法（接收一个用户id，返回用户名称）
    string GetName(1:i32 id)
}
