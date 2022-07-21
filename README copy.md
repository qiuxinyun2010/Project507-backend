# Project507-backend
## Feature
1. 基础功能
 - 用户模块：注册/登录功能；邮箱验证码；用户关注/拉黑功能；
 - 文章模块：用户发表动态；获取用户动态；点赞收藏动态；使用redis缓存热点数据；
 - 标签模块：添加、修改和删除标签；标签检索用户动态；
 - 评论模块：添加用户评论；点赞评论；回复评论；热门评论；
 - 搜索模块：关键词自动补全；历史搜索记录（前端缓存）;

2. 消息功能
 - 即时通信：基于WebSocket协议, 实现用户私信；聊天室群聊；心跳检测；
 - 消息推送：使用redis中间件，实现系统消息的订阅和发布；
 - 历史消息：Mysql保存离线时的消息和私信；定时任务清理过期的历史消息；未读消息计数；
 - 最近会话：使用有序集合缓存用户最近会话列表

3. 后台管理
 - 系统日志：日志文件保存；文件超出大小自动分片；定时清理过期日志；
 - 身份认证：使用JWT中间件实现用户身份认证；
 - 接口管理：基于redis缓存，使用延迟队列，限制IP地址调用特定接口频率、调用间隔、错误次数;

4. 拓展功能 `(TODO)`
 - YOLOv7目标检测算法，自动生成图片标签；通过gRPC调用算法接口
 - 用户相关推荐；
 - 人脸检测和识别；

5. 音视频服务
 - 基于WebRTC