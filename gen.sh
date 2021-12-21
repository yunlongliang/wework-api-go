#/bash
unzip go-client-generated.zip
mv go-client wework-api-go
cd wework-api-go
echo "go mod init"
go mod init github.com/yunlongliang/wework-api-go
echo "go mod tidy"
go mod tidy


echo "更新文件"
cp api_agent.go api_agent.temp

sed -e "19i\\ 
 \"github.com/antihax/optional\"" api_agent.temp > api_agent.go
rm api_agent.temp

cp api_appchat.go api_appchat.temp
sed -e "19i\\ 
 \"github.com/antihax/optional\"" api_appchat.temp > api_appchat.go
rm api_appchat.temp

cp api_groupchat.go api_groupchat.temp
sed -e "19i\\ 
 \"github.com/antihax/optional\"" api_groupchat.temp > api_groupchat.go
 rm api_groupchat.temp

cp api_linkedcorp.go api_linkedcorp.temp
sed -e "19i\\ 
 \"github.com/antihax/optional\"" api_linkedcorp.temp > api_linkedcorp.go
 rm api_linkedcorp.temp

cp api_media.go api_media.temp
sed -e "19i\\ 
 \"github.com/antihax/optional\"" api_media.temp > api_media.go
 rm api_media.temp

cp api_menu.go api_menu.temp
sed -e "19i\\ 
 \"github.com/antihax/optional\"" api_menu.temp > api_menu.go
 rm api_menu.temp

cp api_message.go api_message.temp
sed -e "19i\\ 
 \"github.com/antihax/optional\"" api_message.temp > api_message.go
rm api_message.temp

cp api_service.go api_service.temp
sed -e "19i\\ 
 \"github.com/antihax/optional\"" api_service.temp > api_service.go
rm api_service.temp


echo "git init"
git init
git add .
git remote add origin git@github.com:yunlongliang/wework-api-go.git
git commit -m 'wework-api-go'
git push --set-upstream origin master --force
git tag $1
git push --tag







