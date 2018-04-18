FROM registry.cn-beijing.aliyuncs.com/jsyn/cu-golang-bee:v1

ENV APP_DIR $GOPATH/src/beegoDemo
RUN mkdir -p $APP_DIR
RUN ln -sf /usr/share/zoneinfo/Asia/Shanghai  /etc/localtime
ADD . $APP_DIR
RUN cd $APP_DIR && go build 

EXPOSE 8080
ENTRYPOINT (cd $APP_DIR && bee run)
