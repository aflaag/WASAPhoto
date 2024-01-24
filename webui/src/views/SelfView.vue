<script setup>
	import CommentBox from "../components/CommentBox.vue";
	import ErrorMsg from "../components/ErrorMsg.vue";
	import SuccessMsg from "../components/SuccessMsg.vue";
</script>

<script>
    export default {
		components: { CommentBox, ErrorMsg, SuccessMsg },
        data: function() {
            return {
                errormsg: null,
				successmsg: null,
                loading: false,

				token: localStorage.getItem("token"),
				uname: localStorage.getItem("uname"),

				comments: {},

				show_likes: false,
				likes: null,

				modal: null,

                photos: null,
                empty_photos: true,

                photo_count: 0,
                followers_count: 0,
                following_count: 0,

                show_followers: false,
                show_following: false,

				newUsername: "",

				show_change: false,
				show_change_confirm: false,
            }
        },
        methods: {
			async getPhotoComments(photo) {
				try {
					let response = await this.$axios.get("/user/" + this.uname + "/photos/" + photo.id + "/comments", {
						headers: {
							Authorization: "Bearer " + this.token,
						}
					});

					this.comments = response.data;

					this.modal = new bootstrap.Modal(document.getElementById('logviewer'));
					this.modal.show();
				} catch (e) {
					if (e.response && e.response.status === 500) {
						this.errormsg = "Something went wrong while trying to fetch the photo's comments.";
					} else if (e.response && e.response.status == 401) {
						this.errormsg = "Forbidden access";

						this.$router.replace({path: "/404"});
					} else if (e.response && e.response.status == 404) {
						this.errormsg = "Page not found";

						this.$router.replace({path: "/404"});
					} else {
						this.errormsg = e.toString();
					}
				}
			},
			async updateLike(photo) {
				this.successmsg = null;

				if (!photo.like_status) {
					try {
						let _ = await this.$axios.put("/user/" + photo.user.username + "/photos/" + photo.id + "/likes/" + this.uname, {}, {
							headers: {
								Authorization: "Bearer " + this.token,
							}
						});

						photo.like_count += 1;
					} catch (e) {
						if (e.response && e.response.status === 500) {
							this.errormsg = "Something went wrong while trying to register the like.";
						} else if (e.response && e.response.status == 401) {
							this.errormsg = "Forbidden access";

							this.$router.replace({path: "/404"});
						} else if (e.response && e.response.status == 404) {
							this.errormsg = "Page not found";

							this.$router.replace({path: "/404"});
						} else {
							this.errormsg = e.toString();
						}
					}
				} else {
					try {
						let _ = await this.$axios.delete("/user/" + photo.user.username + "/photos/" + photo.id + "/likes/" + this.uname, {
							headers: {
								Authorization: "Bearer " + this.token,
							}
						});

						photo.like_count -= 1;
					} catch (e) {
						if (e.response && e.response.status === 500) {
							this.errormsg = "Something went wrong while trying to remove the like.";
						} else if (e.response && e.response.status == 401) {
							this.errormsg = "Forbidden access";

							this.$router.replace({path: "/404"});
						} else if (e.response && e.response.status == 404) {
							this.errormsg = "Page not found";

							this.$router.replace({path: "/404"});
						} else {
							this.errormsg = e.toString();
						}
					}
				}

				photo.like_status = !photo.like_status;
			},
			async getPhotoLikes(photo) {
				this.successmsg = null;

				try {
					let response = await this.$axios.get("/user/" + photo.user.username + "/photos/" + photo.id + "/likes", {
						headers: {
							Authorization: "Bearer " + this.token,
						}
					});

					this.likes = response.data;

					this.show_likes = true;

					if (this.likes.users.length > 0) {
						this.empty_likes = false;
					}
				} catch (e) {
					if (e.response && e.response.status === 500) {
						this.errormsg = "Something went wrong while trying to retrieve likes.";
					} else if (e.response && e.response.status == 401) {
						this.errormsg = "Forbidden access";

						this.$router.replace({path: "/404"});
					} else if (e.response && e.response.status == 404) {
						this.errormsg = "Page not found";

						this.$router.replace({path: "/404"});
					} else {
						this.errormsg = e.toString();
					}
				}
			},
			async home() {
				this.successmsg = null;

                this.$router.push({path: "/user/" + this.uname + "/stream"});
			},
            async getProfileInfo() {
				this.successmsg = null;

                try {
					let response = await this.$axios.get("/user/" + this.uname, {
						headers: {
							Authorization: "Bearer " + this.token,
						}
					});

                    this.show_follow = !response.data.follow_status;
                    this.show_ban = !response.data.ban_status;

                    this.photos = response.data.photos;

					if (this.photos.length > 0) {
						this.empty_photos = false;
					}

                    this.photo_count = response.data.photo_count;
                    this.followers_count = response.data.followers_count;
                    this.following_count = response.data.following_count;
				} catch (e) {
					if (e.response && e.response.status === 500) {
						this.errormsg = "Something went wrong while trying to retrieve profile information.";
					} else if (e.response && e.response.status == 401) {
						this.errormsg = "Forbidden access";

						this.$router.replace({path: "/404"});
					} else {
						this.errormsg = e.toString();
					}
				}
            },
            async getFollowers() {
				this.successmsg = null;

                try {
					let response = await this.$axios.get("/user/" + this.uname + "/followers", {
						headers: {
							Authorization: "Bearer " + this.token,
						}
					});

                    this.followers = response.data;

                    this.show_followers = true;
				} catch (e) {
					if (e.response && e.response.status === 500) {
						this.errormsg = "Something went wrong while trying to retrieve the followers.";
					} else if (e.response && e.response.status == 401) {
						this.errormsg = "Forbidden access";

						this.$router.replace({path: "/404"});
					} else {
						this.errormsg = e.toString();
					}
				}
            },
            async getFollowing() {
				this.successmsg = null;

                try {
					let response = await this.$axios.get("/user/" + this.uname + "/following", {
						headers: {
							Authorization: "Bearer " + this.token,
						}
					});

                    this.following = response.data;

                    this.show_following = true;
				} catch (e) {
					if (e.response && e.response.status === 500) {
						this.errormsg = "Something went wrong while trying to retrieve the following.";
					} else if (e.response && e.response.status == 401) {
						this.errormsg = "Forbidden access";

						this.$router.replace({path: "/404"});
					} else {
						this.errormsg = e.toString();
					}
				}
            },
			async doUploadRequest(photoUrl) {
				try {
					if (photoUrl != null) {
						return await this.$axios.post("/user/" + this.uname + "/upload", {
							url: photoUrl,
						}, {
							headers: {
								Authorization: "Bearer " + this.token,
							}
						});
					}	
				} catch (e) {
					if (e.response && e.response.status === 500) {
						this.errormsg = "Something went wrong while trying to upload the photo.";
					} else if (e.response && e.response.status == 401) {
						this.errormsg = "Forbidden access";

						this.$router.replace({path: "/404"});
					} else {
						this.errormsg = e.toString();
					}
				}
			},
			async uploadPhoto() {
				try {
					const file = this.$refs.imageInput.files[0];

					if (file) {
						const reader = new FileReader();

						let photoUrl = null;

						reader.onload = (e) => {
							photoUrl = e.target.result;

							this.doUploadRequest(photoUrl)
								.then(response => {
									this.photos.unshift(response.data);

									if (this.photos.length > 0) {
										this.empty_photos = false;
									}

									this.photo_count += 1;

									this.successmsg = "Photo uploaded correctly!";
								})
						};

						reader.readAsDataURL(file);
					}
				} catch (e) {
					if (e.response && e.response.status === 500) {
						this.errormsg = "Something went wrong while trying to upload the photo.";
					} else if (e.response && e.response.status == 401) {
						this.errormsg = "Forbidden access";

						this.$router.replace({path: "/404"});
					} else {
						this.errormsg = e.toString();
					}
				}
			},
			async deletePhoto(photo) {
				this.successmsg = null;

				try {
					let _ = await this.$axios.delete("/user/" + this.uname + "/photos/" + photo.id, {
						headers: {
							Authorization: "Bearer " + this.token,
						}
					});

                    this.photos = this.photos.filter(function (p) { return p.id !== photo.id});

					this.photo_count -= 1;

					if (this.photos.length === 0) {
						this.empty_photos = true;
					}
				} catch (e) {
					if (e.response && e.response.status === 500) {
						this.errormsg = "Something went wrong while trying to remove the photo.";
					} else if (e.response && e.response.status == 401) {
						this.errormsg = "Forbidden access";

						this.$router.replace({path: "/404"});
					} else if (e.response && e.response.status == 404) {
						this.errormsg = "Forbidden access";

						this.$router.replace({path: "/404"});
					} else {
						this.errormsg = e.toString();
					}
				}
			},
            async changeUsername() {
				this.successmsg = null;

				if (this.newUsername != "") {
					if (this.newUsername === this.uname) {
						this.show_change = false;
						this.show_change_confirm = false;

						this.errormsg = "";
					} else {
						try {

							let response = await this.$axios.put("/user/" + this.uname + "/setusername", {
								username: this.newUsername
							}, {
								headers: {
									Authorization: "Bearer " + this.token,
								}
							});

							let usernameResponse = response.data.username;

							if (usernameResponse === this.uname) {
								this.errormsg = "This username is already taken."
							} else {
								this.uname = usernameResponse;
								localStorage.setItem("uname", usernameResponse);

								this.show_change = false;
								this.show_change_confirm = false;

								this.errormsg = "";
							}
						} catch (e) {
							if (e.response && e.response.status === 500) {
								this.errormsg = "Something went wrong while trying to register the ban.";
							} else if (e.response && e.response.status == 401) {
								this.errormsg = "Forbidden access";

						this.$router.replace({path: "/404"});
							} else {
								this.errormsg = e.toString();
							}
						}
					}
				}
            },
        },
        mounted() {
            this.getProfileInfo()
        }
    }
</script>

<template>
    <div class="everything">
        <div class="header-div">
            <div class="username-div">
                <div style="display: flex; justify-content: center;">
                    <p v-if="!this.show_change" class="header">{{this.uname}}</p>

					<div v-if="this.show_change" style="display: flex; justify-content: center; margin: auto; margin-bottom: 15px; margin-left: -25px">
	                    <input v-model="this.newUsername" class="comment-bar" placeholder="Change it!" style=" font-weight: 800; font-size: 500%; color: #485696; text-align: center;">
					</div>
                </div>

                <div class="option-buttons-div">
					<div style="background-color: #485696; width: 341px; height: 59px; border-radius: 15px;">
						<label for="image" class="btn" style="color: #e7e7e7; font-size: 190%; width: 341px; height: 59px;">Upload photo</label>
						<input type="file" id="image" name="image" accept="image/*" required class="form-input" ref="imageInput" @change="uploadPhoto" style="visibility: hidden;">
					</div>

                    <button v-if="!this.show_change_confirm" @click="this.show_change = true; this.show_change_confirm = true;" class="button">
                        <img class="ban-icon" src="/assets/username.svg">
                    </button>

                    <button v-if="this.show_change_confirm" @click="changeUsername" class="button">
                        <img class="ban-icon" src="/assets/change.svg">
                    </button>
                </div>
            </div>

            <div class="numbers-div">
                <div class="numbers-block-div">
                    <p class="numbers-block-numbers">{{this.photo_count}}</p>
                    <p style="margin-top: 3%">Photos</p>
                </div>

                <button @click="getFollowers" class="button">
                    <div class="numbers-block-div">
                        <p class="numbers-block-numbers">{{this.followers_count}}</p>
                        <p style="margin-top: 3%">Followers</p>
                    </div>
                </button>

                <button @click="getFollowing" class="button">
                    <div class="numbers-block-div">
                        <p class="numbers-block-numbers">{{this.following_count}}</p>
                        <p style="margin-top: 3%">Following</p>
                    </div>
                </button>
            </div>

            <button @click="home" class="button">
                <img class="user-icon" src="/assets/home.svg">
            </button>
        </div>

		<SuccessMsg v-if="successmsg" :msg="successmsg"></SuccessMsg>
		<ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>

		<div v-if="!this.empty_photos" class="horizontal-scroll-panel">
			<div class="post-card" v-for="photo in this.photos" :key="photo.id">
				<div class="post-card-header" style="margin-top: 0px">
					<RouterLink :to="'/user/' + photo.user.username" class="nav-link" style="margin-left: 20px; margin-top: 6px; height: 80px;">
						<p class="post-card-username">{{this.uname}}</p>
					</RouterLink>

					<button @click="deletePhoto(photo)" class="button" style="display: flex; width: 50px; margin-left: 2%; margin-top: -2%">
						<img class="delete-comment" src="/assets/delete.svg"/>
					</button>
				</div>

				<div class="post-photo-div">
					<div class="post-photo-bg"></div>
					<img :src="photo.url" class="post-photo-img">
				</div>

				<div class="post-card-footer" style="margin-top: 7px">
					<button @click="updateLike(photo);" class="button" style="margin-bottom: 60px; margin-left: 20px">
						<div class="post-photo-utils" style="margin-right: 10px;">
							<img v-if="!photo.like_status" src="/assets/like-not-liked.svg"/>
							<img v-if="photo.like_status" src="/assets/like-liked.svg"/>
						</div>
					</button>

					<button @click="getPhotoLikes(photo)" class="button" style="margin: 5px 20px 0px 10px;">
						<p>{{photo.like_count}}</p>
					</button>

					<button @click="getPhotoComments(photo)" class="button" style="margin-bottom: 45px; margin-right: 8px">
						<div class="post-photo-utils">
							<img src="/assets/comment.svg"/>
						</div>
					</button>

					<p>{{photo.comment_count}}</p>

					<CommentBox id="logviewer" :comments="this.comments" :photo="photo" :modal="this.modal"></CommentBox>
				</div>
			</div>
		</div>

		<div v-if="this.empty_photos" class="horizontal-scroll-panel">
			<div style="display: flex; justify-content: center; margin-top: 13%">
				<p style="color: #485696; font-size: 300%">Find new users and follow your friends!</p>
			</div>
		</div>
    </div>

	<div v-if="this.show_likes" class="overlay" style="margin-top: -214px">
		<div class="comment-box">
			<button class="button" @click="this.show_likes = false;" style="display:flex">
				<img class="cross" src="/assets/cross.svg"/>
			</button>
						
			<div class="search-scroll-panel">
				<div v-for="like in this.likes.users" :key="like.id">
					<div class="comment">
						<div class="comment-header">
							<div class="comment-op">
								<RouterLink :to="'/user/' + like.username" class="nav-link">
									<p>{{like.username}}</p>
								</RouterLink>
							</div>
						</div>

						<div class="heightless-line"></div>
					</div>
				</div>
			</div>
		</div>
	</div>

    <div v-if="this.show_followers" class="overlay" style="margin-top: -214px">
		<div class="comment-box">
			<button class="button" @click="this.show_followers = false;" style="display:flex">
				<img class="cross" src="/assets/cross.svg"/>
			</button>
						
			<div class="search-scroll-panel">
				<div v-for="follower in this.followers.users" :key="follower.id">
					<div class="comment">
						<div class="comment-header">
							<div class="comment-op">
								<RouterLink @click="this.show_followers = false;" :to="'/user/' + follower.username" class="nav-link">
									<p>{{follower.username}}</p>
								</RouterLink>
							</div>
						</div>

						<div class="heightless-line"></div>
					</div>
				</div>
			</div>
		</div>
	</div>

    <div v-if="this.show_following" class="overlay" style="margin-top: -214px">
		<div class="comment-box">
			<button class="button" @click="this.show_following = false;" style="display:flex">
				<img class="cross" src="/assets/cross.svg"/>
			</button>
						
			<div class="search-scroll-panel">
				<div v-for="following in this.following.users" :key="following.id">
					<div class="comment">
						<div class="comment-header">
							<div class="comment-op">
								<RouterLink @click="this.show_following = false;" :to="'/user/' + following.username" class="nav-link">
									<p>{{following.username}}</p>
								</RouterLink>
							</div>
						</div>

						<div class="heightless-line"></div>
					</div>
				</div>
			</div>
		</div>
	</div>
</template>

<style>
    .everything {
        display: flex;
        justify-content: center;
        align-items: center;
    }

	.delete-photo {
		width: 10%;
		height: 10%;

		margin-top: 3%;
	}

    .numbers-div {
        display: flex;
        justify-content: space-between;

        margin-top: 2%;

        width: 33%;

        font-weight: 200;
        font-size: 50px;

        color: #485696;
    }

    .numbers-block-div {
        display: flex;
        /* justify-content: center; */
        text-align: center;
        flex-direction: column;
    }

    .numbers-block-numbers {
        font-weight: 600;
        margin-bottom: -10%
    }

    .username-div {
        width: 41%;
    }

    .header {
        font-weight: 800;
        font-size: 500%;
        color: #485696;
    }

    .option-buttons-div {
        /* width: 60%; */
        display: flex;
        justify-content: space-between;

		/* border: solid 3px green; */
    }

	.horizontal-scroll-panel {
		position: absolute;

		/* border: solid 3px red; */

		top: 30%;
		left: 2%;

		width: 96%;
		height: 60%;

		overflow: auto;
		white-space: nowrap;
	}

	.post-card {
		display: inline-block;
		text-align: center;

		width: 25%;
		height: 100%;

		margin-right: 3%;

		border: solid 4px #485696;
		border-radius: 30px;

        color: #485696;

		font-size: 300%;
	}

	.post-photo-div {
		width: 100%;
		height: 70%;

		margin-top: -2%;

		/* border: solid 3px blue; */

		z-index: 1;
	}

	.post-photo-bg {
		width: 100%;
		height: 100%;
		/* object-fit: fill;
		overflow: hidden; */
		background-color: #c8cfe8;

		z-index: 2;
	}

	.post-photo-img {
		width: 100%;
		height: 100%;

		object-fit: contain;
		overflow: hidden;

		z-index: 3;

		/* border: solid 3px blue; */

		/* margin-left: -100%; */
		margin-top: -100%;
	}

	.post-card-footer {
		display: flex;
		font-weight: 800;

		/* border: solid 1px green; */
		height: 14%;
	}

	.post-photo-utils {
		/* border: solid 1px green; */
		height: 90%;

		margin-left: 3%;
	}

	.post-card-footer p {
		margin-top: 1.5%;
		margin-right: 4%;
		margin-left: 2%;
	}

	.post-card-header {
		display: flex;
		justify-content: space-between;
		margin-top: 2%;

		width: 94%;
	}

	.post-card-header p {
		margin-left: 6%;
	}

	::placeholder {
		text-align: center;
	}
</style>