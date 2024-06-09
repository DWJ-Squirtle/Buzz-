<template>
    <div>
        <div v-clickoutside="hideReplyBtn" @click="inputFocus" class="my-reply">
            <el-avatar class="header-img" :size="40" :src="myHeader"></el-avatar>
            <div class="reply-info" >
                <div 
                tabindex="0" 
                contenteditable="true" 
                id="replyInput" 
                spellcheck="false" 
                placeholder="输入评论..." 
                class="reply-input" 
                @focus="showReplyBtn"  
                @input="onDivInput($event)"
                >
                </div>
            </div>
            <div class="reply-btn-box" v-show="btnShow">
                <el-button class="reply-btn" size="medium" @click="sendComment()" type="primary">发表评论</el-button>
            </div>
			<div class="right-buttons">
            <el-button class="other-btn" size="medium" @click="takegood()" type="success">好评</el-button>
            <el-button class="other-btn" size="medium" @click="takebad()"  type="info">差评</el-button>
        </div>
        </div>
        <div v-for="(item,i) in comments" :key="i" class="author-title reply-father" @click="toggleDeleteBtn(i)">
            <el-avatar class="header-img" :size="40" :src="item.headImg"></el-avatar>
            <div class="author-info">
                <span class="author-name">{{item.name}}</span>
                <span class="author-time">{{formattedCreateTime(item.create_time)}}</span>
            </div>
            <div class="talk-box">
                <p>
                    <span class="reply">{{item.content}}</span>
                </p>
            </div>
			<el-button v-if="showDeleteBtn[i]" lass="delete-btn" size="mini"  @click="deleteComment(item.commentIDstr)">删除</el-button>
        </div>
    </div>
</template>

<script>
//import JSONbig from 'json-bigint';
const clickoutside = {
    // 初始化指令
    bind(el,binding) {
    function documentHandler(e) {
    // 这里判断点击的元素是否是本身，是本身，则返回
        if (el.contains(e.target)) {
            return false;
        }
    // 判断指令中是否绑定了函数
        if (binding.expression) {
            // 如果绑定了函数 则调用那个函数，此处binding.value就是handleClose方法
            binding.value(e);
        }
    }
    // 给当前元素绑定个私有变量，方便在unbind中可以解除事件监听
    el.vueClickOutside = documentHandler;
    document.addEventListener('click', documentHandler);
    },
    update() {},
    unbind(el) {
    // 解除事件监听
    document.removeEventListener('click', el.vueClickOutside);
    delete el.vueClickOutside;
  },
};
export default {
    name:'Comment',
	props: {
		sourceId: {
			type: Number,
			require: false,
			default: -1
		}
	},
    data(){
        return{
			btnShow: false,
            index:'0',
			replyComment:'',
            myName:'Lana Del Rey',
            myHeader:'https://n.sinaimg.cn/sinakd20114/0/w1152h2048/20200529/3b22-iufmpmn9683169.jpg',
            myId:19870621,
            to:'',
            toId:-1,
			comments: [],
			showDeleteBtn: []
		}
	},
    directives: {clickoutside},
	isgood:0,
    methods: {
		toggleDeleteBtn(index) {
			this.showDeleteBtn = this.showDeleteBtn.map(() => false);
			this.$set(this.showDeleteBtn, index, true);
			},
		formattedCreateTime(time) {
			var date = new Date(time);
			var year = date.getFullYear();
			var month = ('0' + (date.getMonth() + 1)).slice(-2);
			var day = ('0' + date.getDate()).slice(-2);
			var formattedDate = year + '-' + month + '-' + day;
			return formattedDate;
			},
		takegood(){
			const id=parseInt(this.$route.params.id,10);
			this.$axios.get(`/goodcomment?id=${id}`)
			.then(response => {
				this.comments=response.data;
				this.isgood=1;
			})
			.catch(error =>{
				console.log("fail",error)
			})
		},
		takebad(){
			const id=parseInt(this.$route.params.id);
			this.$axios.get(`/badcomment?id=${id}`)
			.then(response => {
				this.comments=response.data;
				this.isgood=2;
				})
			.catch(error =>{
				console.error("fail",error)
				})
			},
		fetchComments() {
			const id = parseInt(this.$route.params.id, 10);
            this.$axios.get(`/comment?id=${id}`)
            .then(response => {
                this.comments =response.data;
				console.log(response.data);
             })
            .catch(error => {
				this.$message.error('fail-get');
                console.error('获取评论失败:', error);
            });
        },
		deleteComment(comment_id){
			this.$axios.get(`/delete_comment?comment_id=${comment_id}`)
			.then((response) => {
				this.$message.error('删除成功',response.data);
				if(this.isgood==0){
					this.fetchComments();
				}else if(this.isgood==1){
					//this.takegood();
					this.fetchComments();
				}else{
					//this.takebad();
                    this.fetchComments();

				}
			})
			.catch(error =>{
				console.log('删除失败',error);
			});
		},
		sendComment() {
			const id = parseInt(this.$route.params.id, 10);
            if (!this.replyComment) {
                this.$message({
                showClose: true,
                type: 'warning',
                message: '评论不能为空'
                });
            } else {
                this.$axios({
					method: "post",
					url: "/comment",
					data:{
						content: this.replyComment,
						post_id: id,
						parent_id: 12,
						name: "Alice",
						headImg: "https://n.sinaimg.cn/sinakd20114/0/w1152h2048/20200529/3b22-iufmpmn9683169.jpg"
					}
				})
                .then(response => {
					if(response.code==1000){
						this.replyComment="";
						this.fetchComments();
					}else{
						this.$message.error('fale-insert',response.code);
						console.log(response.code);
					}
					})
                .catch(error => {
                    console.log(error);
                    this.$message.error('fail');
                });
            }
        },
        inputFocus(){
            var replyInput = document.getElementById('replyInput');
            replyInput.style.padding= "8px 8px"
            replyInput.style.border ="2px solid blue"
            replyInput.focus()
        },  
        showReplyBtn(){
            this.btnShow = true
        },
        hideReplyBtn(){
            this.btnShow = false
            var replyInput = document.getElementById('replyInput');
			replyInput.style.padding= "10px"
            replyInput.style.border ="none"
        },
        showReplyInput(i,name,id){
            this.comments[this.index].inputShow = false
            this.index =i
            this.comments[i].inputShow = true
            this.to = name
            this.toId = id
        },
        _inputShow(i){
            return this.comments[i].inputShow 
        },
        onDivInput: function(e) {
            this.replyComment = e.target.innerHTML;
        },
        dateStr(date){
            //获取js 时间戳
            var time=new Date().getTime();
            //去掉 js 时间戳后三位，与php 时间戳保持一致
            time=parseInt((time-date)/1000);
            //存储转换值 
            var s;
            if(time<60*10){//十分钟内
                return '刚刚';
            }else if((time<60*60)&&(time>=60*10)){
                //超过十分钟少于1小时
                s = Math.floor(time/60);
                return  s+"分钟前";
            }else if((time<60*60*24)&&(time>=60*60)){ 
                //超过1小时少于24小时
                s = Math.floor(time/60/60);
                return  s+"小时前";
            }else if((time<60*60*24*30)&&(time>=60*60*24)){ 
                //超过1天少于30天内
                s = Math.floor(time/60/60/24);
                return s+"天前";
            }else{ 
                //超过30天ddd
                date= new Date(parseInt(date));
                return date.getFullYear()+"/"+(date.getMonth()+1)+"/"+date.getDate();
            }
        }
    },
	mounted() {
			this.fetchComments();
    },
}
</script>

<style scoped>
.my-reply {
    padding: 10px;
	margin-top: 100px;
    background-color: #fafbfc;
}
.my-reply .header-img {
    display: inline-block;
    vertical-align: top;
}
.my-reply .reply-info {
    display: inline-block;
    margin-left: 5px;
    width: 90%;
}
@media screen and (max-width: 1200px) {
    .my-reply .reply-info {
        width: 80%;
    }
}
.my-reply .reply-info .reply-input {
    min-height: 20px;
    line-height: 22px;
    padding: 10px 10px;
    color: #ccc;
    background-color: #fff;
    border-radius: 5px;
}
.my-reply .reply-info .reply-input:empty:before {
    content: attr(placeholder);
}
.my-reply .reply-info .reply-input:focus:before {
    content: none;
}
.my-reply .reply-info .reply-input:focus {
    padding: 8px 8px;
    border: 2px solid blue;
    box-shadow: none;
    outline: none;
}
.my-reply .reply-btn-box {
    height: 25px;
    margin: 10px 0;
}
.my-reply .reply-btn-box .reply-btn{
    position: relative;
    float: right;
    margin-right: 15px;
}
.delete-btn{
	position: relative;
	top: 0;
	float: right;
	right: 15px;
}

.my-comment-reply {
    margin-left: 50px;
}
.my-comment-reply .reply-input {
    width: flex;
}
.author-title:not(:last-child) {
    border-bottom: 1px solid rgba(178,186,194,.3);
}
.author-title {
    padding: 10px;
}
.author-title .header-img {
    display: inline-block;
    vertical-align: top;
}
.author-title .author-info {
    display: inline-block;
    margin-left: 5px;
    width: 60%;
    height: 40px;
    line-height: 20px;
}
.author-title .author-info > span {
    display: block;
    cursor: pointer;
    overflow: hidden;
    white-space: nowrap;
    text-overflow: ellipsis;
}
.author-title .author-info .author-name {
    color: #000;
    font-size: 18px;
    font-weight: bold;
}
.author-title .author-info .author-time {
    font-size: 14px;
}
.author-title .icon-btn {
    width: 30%;
    padding: 0 !important;
    float: right;
}
@media screen and (max-width: 1200px) {
    .author-title .icon-btn {
        width: 20%;
        padding: 7px;
    }
}
.author-title .icon-btn > span {
    cursor: pointer;
}
.author-title .icon-btn .iconfont {
    margin: 0 5px;
}
.author-title .talk-box {
    margin: 0 50px;
}
.author-title .talk-box > p {
    margin: 0;
}
.author-title .talk-box .reply {
    font-size: 16px;
    color: #000;
}
.author-title .reply-box {
    margin: 10px 0 0 50px;
    background-color: #efefef;
}
.right-buttons {
    display: inline-block;
    vertical-align: top;
    margin-left: 5px; /* 调整按钮容器与文本框之间的间距 */
}

.other-btn {
    display: inline-block;
    vertical-align: top;
    margin-left: 5px; /* 调整按钮与其他按钮之间的间距 */
}

</style>

