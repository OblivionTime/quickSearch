<template>
    <div class="header">
        <div class="items">
            <button class="item" :index="index" @contextmenu.prevent="groupopenMenu"
                :class="item.value == active_id ? 'item-active' : ''" v-for="item, index in CategoryList" :key="item.value"
                @click="(active_id = item.value), loadData()">
                {{ item.label }}
            </button>
        </div>
    </div>
    <div class="operation">
        <div class="left">
            <el-button type="primary" class="btn" @click="ToAddCategory">添加分组</el-button>
        </div>
        <div class="right">
            <el-input v-model="searchInput" class="search" clearable placeholder="搜索" :suffix-icon="Search"
                @input="loadData" />
        </div>

    </div>
    <div class="container" @dragover.prevent @drop.prevent="handleDrop">

        <div class="item" v-for="item, index in FileList" :index="index" :key="item.id" @contextmenu.prevent="fileopenMenu"
            @click="OpenFile(item)">
            <div class="img">
                <img :src="item.file_type == '文件夹' ? Folder : Files" alt="">
            </div>
            <div class="item-content">
                <div class="title">
                    <span v-html="highlightedText(item.name)"></span>
                    <div class="tag">
                        <el-tag v-if="item.file_type == '文件夹'">文件夹</el-tag>
                        <el-tag v-else type="success"> 文件</el-tag>
                    </div>
                </div>
                <div class="desc">
                    <span v-html="highlightedText(item.file_path)"></span>
                </div>
            </div>

        </div>
    </div>
    <el-dialog v-model="addFileDialog" title="文件/文件夹上传" width="50%" @beforeClose="addFileDialog = false">
        <el-form label-position="top" label-width="100px" :model="FileForData" style="max-width: 460px">
            <el-form-item label="名称">
                <el-input v-model="FileForData.name" />
            </el-form-item>
            <el-form-item label="文件/文件夹路径">
                <el-input v-model="FileForData.file_path" readonly />
            </el-form-item>
            <el-form-item label="分组">
                <el-select v-model="FileForData.column_id" style="width:100%">
                    <el-option :label="item.label" :value="item.value" v-for="item in CategoryList.slice(1)"
                        :key="item.value" />
                </el-select>
            </el-form-item>

        </el-form>
        <template #footer>
            <span class="dialog-footer">
                <el-button @click="addFileDialog = false">取消</el-button>
                <el-button type="primary" @click="UploadFile">
                    上传
                </el-button>
            </span>
        </template>
    </el-dialog>
    <rightClickMenu ref="rightClickMenuRef" :RightMenu="GroupRightMenu"></rightClickMenu>
    <rightClickMenu ref="rightClickMenuRef2" :RightMenu="FileRightMenu"></rightClickMenu>
</template>

<script setup>
import { Search } from '@element-plus/icons-vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import Files from '@/assets/files.svg';
import Folder from '@/assets/Folder.svg';
import EditIcon from '@/assets/edit.svg';
import DeleteIcon from '@/assets/delete.svg';
import { getInfo, AddCategory, UpdateCategory, DeleteCategory, AddFileInfo, UpdateFileInfo, DeleteFileInfo } from '@/api/info.js';
import rightClickMenu from '@/components/rightClickMenu/rightClickMenu.vue';
const FileList = ref([])
const CategoryList = ref([])
const searchInput = ref("")
const active_id = ref(0)
const loadData = () => {
    getInfo({ id: active_id.value, name: searchInput.value })
        .then((res) => {
            CategoryList.value = [{
                label: "全部",
                value: 0
            }]
            for (const item of res.data.categoryList) {
                CategoryList.value.push({
                    label: item.name,
                    value: item.id
                })
            }
            FileList.value = res.data.fileList
        }).catch((err) => {
            setTimeout(() => {
                loadData()
            }, 1000);
        });
}
const highlightedText = (text) => {
    if (!searchInput.value) {
        return text
    }

    // 使用正则表达式进行匹配，并添加样式
    const regex = new RegExp(`(${searchInput.value})`, 'gi');
    const highlightedContent = text.replace(regex, '<span style="color: red;">$1</span>');

    return highlightedContent;
}
const { shell } = require('electron');

const OpenFile = (item) => {
    shell.openPath(item.file_path);
}
/**
 * 
 * 文件上传相关
 */
const addFileDialog = ref(false)
const FileForData = ref({

})
const isDirectory = (path) => {
    const fs = require('fs');
    const stats = fs.statSync(path);
    return stats.isDirectory();
}
const handlerCateGory = () => {

    for (const item of CategoryList.value) {
        if (item.value == active_id.value) {
            return
        }
    }
}
const handleDrop = (event) => {
    try {
        if (addFileDialog.value) {
            return
        }
        let column_id = active_id.value
        if (active_id.value == 0) {
            column_id = CategoryList.value[1].value
        }
        FileForData.value = {
            name: event.dataTransfer.files[0].name,
            file_path: event.dataTransfer.files[0].path,
            file_type: isDirectory(event.dataTransfer.files[0].path) ? "文件夹" : "文件",
            column_id: column_id,
        }
        addFileDialog.value = true
    } catch (error) {
        ElMessage.error(' 出现错误,请重启应用!!!')
    }

}
const UploadFile = () => {
    if (FileForData.value.name == "") {
        return ElMessage.error('请输入名称!!')
    }
    AddFileInfo(FileForData.value)
        .then((res) => {
            if (res.code == 0) {
                ElMessage({
                    message: '上传成功!!!',
                    type: 'success',
                })
                addFileDialog.value = false
                loadData()
            } else {
                ElMessage.error(res.msg)
            }
        }).catch((err) => {
            ElMessage.error('上传失败!!')
        });
}
/**
 * 文件相关
 */
const ToAddCategory = () => {
    ElMessageBox.prompt('请输入分组名', '添加分组', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        inputValidator: inputValidator,
    })
        .then(({ value }) => {
            AddCategory({ name: value })
                .then((res) => {
                    if (res.code == 0) {
                        ElMessage({
                            message: '添加成功!!!',
                            type: 'success',
                        })
                        loadData()
                    } else {
                        ElMessage.error('添加失败!!')
                    }
                }).catch((err) => {
                    ElMessage.error('添加失败!!')
                });
        }).catch((err) => {

        });

}
function inputValidator(text) {
    if (text.trim()) {
        return true
    } else {
        return "不能为空"
    }
}
//右键菜单
/**
 * 分组的右键菜单
 */
const GroupRightMenu = [
    {
        name: "重命名",
        icon: EditIcon,
        action: () => {
            if (rightClickMenuRef.value.active_item == 0) {
                return ElMessage.error('不能修改当前分组!!')
            }
            var t = JSON.parse(JSON.stringify(CategoryList.value[rightClickMenuRef.value.active_item]))
            ElMessageBox.prompt('请输入分组名', '修改分组', {
                confirmButtonText: '确定',
                cancelButtonText: '取消',
                inputValue: t.label,
                inputValidator: inputValidator,
            })
                .then(({ value }) => {
                    t.name = value
                    UpdateCategory({ id: t.value, name: t.name })
                        .then((res) => {
                            if (res.code == 0) {
                                ElMessage({
                                    message: '修改成功!!!',
                                    type: 'success',
                                })
                                loadData()
                            } else {
                                ElMessage.error(res.msg)
                            }
                        }).catch((err) => {
                            ElMessage.error('修改失败!!')
                        });
                }).catch((err) => {

                });

        }
    },
    {
        name: "删除分组",
        icon: DeleteIcon,
        action: () => {
            if (rightClickMenuRef.value.active_item == 0) {
                return ElMessage.error('不能删除当前分组!!')
            }
            const id = CategoryList.value[rightClickMenuRef.value.active_item].value
            ElMessageBox.confirm(
                `此操作将永远删除<span style='color:red'>${CategoryList.value[rightClickMenuRef.value.active_item].label}</span>,确定要继续吗?`,
                'Warning',
                {
                    confirmButtonText: '删除',
                    cancelButtonText: '取消',
                    type: 'warning',
                    dangerouslyUseHTMLString: true,
                }
            )
                .then(() => {
                    DeleteCategory({ id: id })
                        .then((res) => {
                            if (res.code == 0) {
                                ElMessage({
                                    message: '修改成功!!!',
                                    type: 'success',
                                })
                                loadData()
                            } else {
                                ElMessage.error(res.msg)
                            }
                        }).catch((err) => {
                            ElMessage.error('修改失败!!')
                        });
                })
                .catch((err) => {
                    console.log(err);
                })

        }
    },
]
const rightClickMenuRef = ref(null);
const groupopenMenu = (e) => {
    rightClickMenuRef.value.openMenu(e)
    rightClickMenuRef2.value.closeMenu()
}
const FileRightMenu = [
    {
        name: "重命名",
        icon: EditIcon,
        action: () => {
            let id = FileList.value[rightClickMenuRef2.value.active_item].id
            ElMessageBox.prompt('请输入名称', '修改名称', {
                confirmButtonText: '确定',
                cancelButtonText: '取消',
                inputValue: FileList.value[rightClickMenuRef2.value.active_item].name,
                inputValidator: inputValidator,
            })
                .then(({ value }) => {
                    UpdateFileInfo({ id, name: value })
                        .then((res) => {
                            if (res.code == 0) {
                                ElMessage({
                                    message: '修改成功!!!',
                                    type: 'success',
                                })
                                loadData()
                            } else {
                                ElMessage.error(res.msg)
                            }
                        }).catch((err) => {
                            ElMessage.error('修改失败!!')
                        });
                }).catch((err) => {

                });
        }
    },
    {
        name: "取消收藏",
        icon: DeleteIcon,
        action: () => {
            const id = FileList.value[rightClickMenuRef2.value.active_item].id
            ElMessageBox.confirm(
                `此操作将永远删除<span style='color:red'>${FileList.value[rightClickMenuRef2.value.active_item].name}</span>,确定要继续吗?`,
                'Warning',
                {
                    confirmButtonText: '删除',
                    cancelButtonText: '取消',
                    type: 'warning',
                    dangerouslyUseHTMLString: true,
                }
            )
                .then(() => {
                    DeleteFileInfo({ id: id })
                        .then((res) => {
                            if (res.code == 0) {
                                ElMessage({
                                    message: '修改成功!!!',
                                    type: 'success',
                                })
                                loadData()
                            } else {
                                ElMessage.error(res.msg)
                            }
                        }).catch((err) => {
                            ElMessage.error('修改失败!!')
                        });
                })
                .catch((err) => {
                    console.log(err);
                })


        }
    },
]
const rightClickMenuRef2 = ref(null);
const fileopenMenu = (e) => {
    rightClickMenuRef2.value.openMenu(e)
    rightClickMenuRef.value.closeMenu()
}
onMounted(() => {
    loadData()

})
</script>

<style lang="scss" scoped>
.header {
    box-sizing: border-box;
    position: relative;
    border-radius: 10px;
    background: white;
    height: 50px;

    .items {
        display: flex;
        align-items: center;
        gap: 10px;
        flex-wrap: wrap;
        box-sizing: border-box;

        .item {

            font-size: 15px;
            cursor: pointer;
            width: 110px;
            height: 50px;
            white-space: nowrap;
            overflow: hidden;
            text-align: center;
            text-overflow: ellipsis;
            transition: all 0.2s ease-in-out;
            background: transparent;
            border-radius: 5px;
            border: 0;

            &-active {
                color: #005EEB !important;
                border: 1px solid #005EEB;
                border-radius: 5px;
            }

            &:hover {
                color: #005EEB;
            }
        }
    }



}

.operation {
    margin: 20px 0;
    display: flex;
    justify-content: space-between;
    align-items: center;

    .btn {
        width: 120px;
        height: 35px;
    }

    .search {
        width: 300px;
        height: 35px;
    }
}

.container {
    padding: 20px;
    border-radius: 10px;
    min-height: 60vh;
    background: white;
    display: grid;
    grid-template-columns: repeat(3, 1fr);
    grid-gap: 20px;

    .item {
        cursor: pointer;
        padding: 20px;
        display: flex;
        box-shadow: 0px 0px 3px rgba($color: #000, $alpha: 0.5);
        gap: 30px;
        height: 100px;
        border-radius: 10px;
        position: relative;
        transition: all 0.2s ease-in-out;

        &:hover {
            box-shadow: 0px 0px 5px rgba($color: #087f5b, $alpha: 0.8);
            transform: scale(1.1);
        }

        .img {
            box-shadow: 0px 0px 5px rgba($color: #000, $alpha: 0.5);
            border-radius: 10px;
            height: 50px;
            width: 50px;
            text-align: center;

            img {
                width: 45px;
                height: 45px;
            }
        }

        .title {
            font-size: 20px;
            width: 180px;
            white-space: nowrap;
            overflow: hidden;
            text-overflow: ellipsis;

            .tag {
                position: absolute;
                right: 20px;
                top: 20px;
            }


        }

        .desc {
            margin-top: 10px;
            font-size: 12px;
            color: #868e96;
            word-break: break-all;
        }
    }
}
</style>