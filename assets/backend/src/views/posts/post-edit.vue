<style lang="less">
    @import './post-edit.less';
    @import '../../styles/simplemde.min.css';
</style>

<template>
    <div class="pit-post-form" style="margin: 20px 0 0 10px">
        <Form :model="myForm" ref="myForm" :rules="myRules" label-position="top">
            <Row>
                <Col span="18" style="padding-right: 20px;">
                <FormItem label="" prop="title">
                    <Input v-model="myForm.title" placeholder="标题" @on-blur="titleBlur"></Input>
                </FormItem>
                <FormItem label="">
                    <markdown-editor id="markdown_editor" v-model="myForm.markdown" :configs="simpleMDEConfigs"
                                     ref="markdownEditor"></markdown-editor>
                </FormItem>
                </Col>
                <Col span="6">
                <FormItem label="分类" prop="categoryId">
                    <Select v-model="myForm.categoryId">
                        <Option v-for="item in categorys" :value="item.id" :key="item.id">{{ item.name }}</Option>
                    </Select>
                </FormItem>
                <FormItem label="别名" prop="slug">
                    <Input type="textarea" v-model="myForm.slug"></Input>
                </FormItem>
                <FormItem label="标签" prop="tags">
                    <Tag v-for="tag in myForm.tags" color="blue" :key="tag" :name="tag" closable
                         @on-close="closeTags">{{ tag }}</Tag>
                    <Input v-model="newTag" @on-enter="addTag" @on-blur="addTag" size="small" placeholder="+" style="width: 50px;"></Input>
                </FormItem>
                <FormItem label="封面图">
                    <Upload v-show="uploadBtnShow"
                            :action="uploadsUrl"
                            :headers="headers"
                            :format="['jpg','jpeg','png','bmp']"
                            :max-size="2048"
                            :show-upload-list="false"
                            :on-progress="uploadsProgress"
                            :on-success="uploadsSuccess"
                            :on-format-error="uploadsFormatError"
                            :on-exceeded-size="uploadsMaxSize"
                            :before-upload="uploadsBeforeUpload">
                        <Button type="ghost" icon="ios-cloud-upload-outline">选择图片</Button>
                    </Upload>
                    <Progress v-show="uploadPercentShow" :percent="uploadPercent" :stroke-width="5">
                        <Icon type="checkmark-circled"></Icon>
                        <span>成功</span>
                    </Progress>
                    <template v-if="myForm.thumb">
                        <div class="show-thumb">
                            <img :src="myForm.thumb" alt=""/>
                            <a @click="clearThumb()" href="javascript:void(0)">移除封面图</a>
                        </div>
                    </template>
                </FormItem>
                <FormItem label="">
                    <Button type="ghost" @click="closeForm()">取 消</Button>
                    <Button type="primary" @click="submitMyForm()">确 定</Button>
                </FormItem>
                </Col>
            </Row>
        </Form>
    </div>
</template>

<script>
    import util from '../../libs/util';
    import inlineAttachment from '../../libs/inline-attachment';
    import markdownEditor from 'vue-simplemde/src/markdown-editor';

    export default {
        components: {
            markdownEditor
        },
        data () {
            return {
                uploadBtnShow: true,
                uploadPercentShow: false,
                uploadPercent: 0,
                editFormLoading: false,
                loading: false,
                myFormTitle: '编辑',
                domain: 'http://',
                newTag: '',
                textareaRow: 15,
                categorys: [],
                markdownPreviews: '',
                uploadsUrl: persimmon.baseUrl + '/backend/files/uploads',
                previewVisible: false,
                showPreview: '',
                myFiles: [],
                myForm: {
                    id: 0,
                    title: '',
                    slug: '',
                    thumb: '',
                    tags: [],
                    categoryId: 0,
                    markdown: ''
                },
                headers: {'X-CSRFToken': persimmon.csrf},
                myRules: {
                    title: [
                        {required: true, type: 'string', message: '请填写文章标题', trigger: 'blur'}
                    ],
                    slug: [
                        {required: true, type: 'string', message: '请填写文章别名', trigger: 'blur'},
                        {pattern: /^[a-zA-Z0-9_-]+$/, message: '只允许英文或者拼音,横杠(-),下划线(_)', trigger: 'blur'}
                    ],
                    categoryId: [
                        {required: true, type: 'integer', message: '请选择分类', trigger: 'blur'}
                    ],
                    markdown: [
                        {required: true, type: 'string', message: '文章内容不能为空', trigger: 'blur'}
                    ]
                },
                simpleMDEConfigs: {
                    placeholder: '日志正文',
                    autofocus: true,
                    autosave: {
                        enabled: true,
                        uniqueId: 'MyPersimmonMarkdown',
                        delay: 1000,
                    },
                    autoDownloadFontAwesome: false,
                    spellChecker: false,
                    styleSelectedText: false,
                    renderingConfig: {
                        singleLineBreaks: true,
                        codeSyntaxHighlighting: true
                    },
                    toolbar: ['bold', 'italic', 'strikethrough', 'heading-1', 'heading-2', 'heading-3', '|', 'code', 'quote', 'unordered-list', 'clean-block', '|', 'link', 'image', {
                        name: 'uploadFile',
                        action: function (editor) {
                            util.uploadFile(editor);
                        },
                        className: 'fa fa-cloud-upload',
                        title: 'upload Files',
                    }, 'table', 'horizontal-rule', '|', 'preview', 'guide']
                }
            };
        },
        computed: {
            simplemde () {
                return this.$refs.markdownEditor.simplemde;
            }
        },
        created () {
            let postId = this.$route.params.id;
            if (postId !== undefined) {
                this.getPost(postId);
            }
            if (this.$route.name === 'add') {
                this.removeItem();
            }
        },
        methods: {
            uploadsBeforeUpload () {
                this.uploadPercentShow = true;
            },
            uploadsFormatError (file) {
                this.$Notice.warning({
                    title: '文件格式错误',
                    desc: '文件 ' + file.name + ' 格式不对, 请选择 jpg 或者 png.'
                });
            },
            uploadsMaxSize (file) {
                this.$Notice.warning({
                    title: '文件过大',
                    desc: '文件  ' + file.name + ' 太大了，请选择 2M 以内的图片.'
                });
            },
            uploadsProgress (event, file, fileList) {
                let that = this;
                that.uploadPercent = event.percent;
                if (event.percent === 100) {
                    setTimeout(function () {
                        that.uploadPercentShow = false;
                    }, 1 * 1000);
                }
            },
            uploadsSuccess (res, file) {
                if (res.status === 200) {
                    this.myForm.thumb = res.item;
                } else {
                    this.myForm.thumb = '';
                }
                this.uploadBtnShow = false;
            },
            clearThumb () {
                this.myForm.thumb = '';
            },
            clearCache () {
                this.removeItem();
                this.myForm.markdown = '';
                //console.log('Key is cleared!');
            },
            closeTags (tag) {
                this.myForm.tags.splice(this.myForm.tags.indexOf(tag), 1);
            },
            addTag () {
                let inputValue = this.newTag;
                if (inputValue) {
                    this.myForm.tags.push(inputValue);
                }
                this.newTag = '';
            },
            titleBlur (event) {
                let that = this;
                if (that.myForm.slug !== '') {
                    return false;
                }
                let query = that.myForm.title;
                if (query.match(/\w+/g) != null) {
                    that.myForm.slug = query;
                }
                if (query == null || query === '') {
                    return false;
                }
                util.ajax.get('/backend/utils/fanyi/' + query).then(function (response) {
                    let res = response.data;
                    if (res.status === 200) {
                        that.myForm.slug = res.item;
                    }
                }).catch(function (error) {
                    console.log(error);
                });
            },
            getPost (id) {
                if (parseInt(id) <= 0) {
                    return false;
                }
                let that = this;
                that.editFormLoading = true;
                that.myFormTitle = '编辑';
                util.ajax.get('/backend/posts/' + id).then(function (response) {
                    let res = response.data;
                    if (res.status === 200) {
                        let item = res.item;
                        let tags = [];
                        that.myForm = item;
                        //tags
                        if (item.tags.length > 0) {
                            for (let index in item.tags) {
                                tags.push(item.tags[index].name);
                            }
                            that.myForm.tags = tags;
                        }
                        //thumb
                        if (item.thumb !== '') {
                            that.imageList = [{url: item.thumb}];
                        }
                    } else {
                        that.$Notice.error({
                            title: '数据获取失败',
                            desc: ''
                        });
                    }
                    that.editFormLoading = false;
                }).catch(function (error) {
                    console.log(error);
                    that.editFormLoading = false;
                });
            },
            submitMyForm () {
                let that = this;
                this.$refs['myForm'].validate((valid) => {
                    if (!valid) {
                        console.log('myForm valid error.');
                        return false;
                    }
                    if (that.myForm.id > 0) {
                        util.ajax.post('/backend/posts/update', that.myForm).then(function (response) {
                            let res = response.data;
                            that.$Notice.open({
                                title: res.status === 200 ? '编辑成功' : '编辑失败',
                                desc: ''
                            });
                            if (res.status === 200) {
                                that.closeForm('myForm');
                            }
                        }).catch(function (error) {
                            console.log(error);
                        });
                    } else {
                        util.ajax.post('/backend/posts/store', that.myForm).then(function (response) {
                            let res = response.data;
                            if (res.status === 200) {
                                that.closeForm('myForm');
                            }
                            that.$Notice.open({
                                title: res.status === 200 ? '新增成功' : '新增失败',
                                desc: ''
                            });
                        }).catch(function (error) {
                            if (error.response) {
                                if (error.response.status !== 200) {
                                    for (let index in error.response.data) {
                                        that.$Notice.warning({
                                            title: '警告',
                                            desc: error.response.data[index][0]
                                        });
                                    }
                                }
                            } else {
                                console.log(error);
                            }
                        });
                    }
                });
            },
            closeForm () {
                this.clearCache();
                this.$refs['myForm'].resetFields();
                this.$router.push({path: '/posts/list'});
            },
            removeItem () {
                localStorage.removeItem('smde_MyPersimmonMarkdown');
            },
            getCategorys () {
                let that = this;
                util.ajax.get('/backend/categories', {
                    params: {
                        rows: 999,
                        page: 1
                    }
                }).then(function (response) {
                    let res = response.data;
                    if (res.status === 200) {
                        that.categorys = res.list.data;
                    } else {
                        that.$Notice.warning({
                            title: '数据获取失败',
                            desc: ''
                        });
                    }
                }).catch(function (error) {
                    console.log(error);
                });
            },
            setDomain () {
                let location = window.location;
                this.domain = location.protocol + '//' + location.host + '/';
            },
            inlineAtta () {
                let that = this;
                inlineAttachment.editors.codemirror4.attach(that.simplemde.codemirror, {
                    uploadUrl: that.uploadsUrl,
                    progressText: '![uploading file...]()',
                    urlText: '![]({filename})',
                    errorText: 'Error uploading file',
                    jsonFieldName: 'file',
                    uploadFieldName: 'file',
                    extraParams: {},
                    extraHeaders: that.headers,
                    onFileUploadResponse: function (xhr) {
                        let result = JSON.parse(xhr.responseText);
                        if (result.status === 200) {
                            let filename = result.item;
                            if (result && filename) {
                                let newValue;
                                if (typeof this.settings.urlText === 'function') {
                                    newValue = this.settings.urlText.call(this, filename, result);
                                } else {
                                    newValue = this.settings.urlText.replace(this.filenameTag, filename);
                                }
                                let text = this.editor.getValue().replace(this.lastValue, newValue);
                                this.editor.setValue(text);
                                this.settings.onFileUploaded.call(this, filename);
                            }
                            return false;
                        }
                    },
                    onFileUploadError: function (data) {
                        that.$Notice.warning({
                            title: '很抱歉，上传失败。请重试。',
                            desc: ''
                        });
                    }
                });
            },
        },
        watch: {
            '$route' (to, from) {
                //监听路由改变
                let postId = this.$route.params.id;
                if (postId !== undefined) {
                    this.getPost(postId);
                }
            },
            'myForm.thumb' (to, form) {
                if (to !== '') {
                    this.uploadBtnShow = false;
                } else {
                    this.uploadBtnShow = true;
                }
            }
        },
        mounted () {
            this.getCategorys();
            this.setDomain();
            //this.simpleMDE();
            this.inlineAtta();
        }
    };
</script>
