<script>
    export default {
        props: ["comments", "photo", "modal"],
        data() {
            return {
                errormsg: null,
                loading: false,

				token: localStorage.getItem("token"),
				uname: localStorage.getItem("uname"),

                comment_body: "",
            }
        },
        methods: {
            async postComment(comments, photo) {
                if (this.comment_body === "") {
                    this.errormsg = "The comment is empty";
                } else {
                    try {
                        let response = await this.$axios.post("/user/" + photo.user.username + "/photos/" + photo.id + "/comment", {
                            user: {
                                id: parseInt(this.token),
                                username: this.uname,
                            },
                            comment_body: this.comment_body,
                        }, {
							headers: {
								Authorization: "Bearer " + this.token,
							}
                        });

                        comments.push(response.data);

                        photo.comment_count += 1;
                    } catch (e) {
                        if (e.response && e.response.status === 500) {
                            this.errormsg = "Something went wrong while trying to post the comment.";
                        } else if (e.response && e.response.status == 401) {
                            this.errormgs = "Forbidden access"
                        } else {
                            this.errormsg = e.toString();
                        }
                    }
                }
            },
            async deleteComment(comments, comment, photo) {
                try {
                    let _ = await this.$axios.delete("/user/" + photo.user.username + "/photos/" + photo.id + "/comments/" + comment.id, {
                        headers: {
                            Authorization: "Bearer " + this.token,
                        }
                    }, {});

                    comments = comments.filter(function (c) { return c.id !== comment.id});

                    photo.comment_count -= 1;
                } catch (e) {
                    if (e.response && e.response.status === 500) {
                        this.errormsg = "Something went wrong while trying to delete the comment.";
                    } else if (e.response && e.response.status == 401) {
                        this.errormgs = "Forbidden access"
                    } else {
                        this.errormsg = e.toString();
                    }
                }
            },
        },
        mounted() {}
    }
</script>

<template>
	<div class="modal modal-xl" tabindex="-1">
		<div class="modal-dialog comment-box" style="margin-top: 170px">
			<div class="modal-content">
				<div class="modal-header">
                    <button class="button" style="display:flex;" data-bs-dismiss="modal">
                        <img class="cross" src="/assets/cross.svg"/>
                    </button>
				</div>
				<div class="modal-body">
					<div class="comment-scroll-panel" style="height: 500px; position: absolute; margin-top: -50px; margin-left: -15px">
						<div class="comment" v-for="comment in comments.comments" :key="comment.id">
                            <div class="comment-header">
								<RouterLink @click="modal.hide()" :to="'/user/' + comment.user.username" class="nav-link">
									<p>{{comment.user.username}}</p>
								</RouterLink>

                                <button v-if="comment.user.username === this.uname" @click="deleteComment(comments.comments, comment, photo)" class="button" style="display: flex; width: 50px; margin-left: 2%; margin-top: -2%">
                                    <img class="delete-comment" src="/assets/delete.svg"/>
                                </button>
                            </div>

                            <div class="comment-text">
                                <p>{{comment.comment_body}}</p>
                            </div>

                            <div class="heightless-line"></div>
						</div>
					</div>
				</div>

                <div class="comment-input-box">
                    <input v-model="comment_body" class="comment-bar" placeholder="Leave a comment!">

                    <button class="button" @click="postComment(comments.comments, photo)">
                        <img class="button-image" src="/assets/arrow.svg" style="width: 70%"/>
                    </button>
                </div>
			</div>
		</div>
	</div>
</template>

<style>
    .modal-content {
		position: relative;

		width: 100%;
		height: 100%;

		border-radius: 30px;

		background-color: #e7e7e7;
        color: black;

		font-size: 40%;
        font-weight: 400;

		z-index: 1;
    }
</style>