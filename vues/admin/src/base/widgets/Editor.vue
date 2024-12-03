<template>
  <div class="g-editor">
    <input
      accept="image/*"
      ref="mfiles"
      type="file"
      style="display: none"
      @change="fileChange"
    />
    <div ref="header" class="header">
      <span class="ql-formats">
        <select class="ql-font"></select>
        <select class="ql-size"></select>
      </span>
      <span class="ql-formats">
        <button class="ql-bold"></button>
        <button class="ql-italic"></button>
        <button class="ql-underline"></button>
        <button class="ql-strike"></button>
      </span>
      <span class="ql-formats">
        <select class="ql-color"></select>
        <select class="ql-background"></select>
      </span>
      <span class="ql-formats">
        <button class="ql-script" value="sub"></button>
        <button class="ql-script" value="super"></button>
      </span>
      <span class="ql-formats">
        <button class="ql-header" value="1"></button>
        <button class="ql-header" value="2"></button>
        <button class="ql-blockquote"></button>
        <button class="ql-code-block"></button>
      </span>
      <span class="ql-formats">
        <button class="ql-list" value="ordered"></button>
        <button class="ql-list" value="bullet"></button>
        <button class="ql-indent" value="-1"></button>
        <button class="ql-indent" value="+1"></button>
      </span>
      <span class="ql-formats">
        <button class="ql-direction" value="rtl"></button>
        <select class="ql-align"></select>
      </span>
      <span class="ql-formats">
        <button class="ql-link"></button>
        <button class="ql-image"></button>
        <button class="ql-video"></button>
        <button class="ql-formula"></button>
      </span>
      <span class="ql-formats">
        <button class="ql-clean"></button>
      </span>
    </div>

    <div class="body" ref="editor_quill">
      <div v-html="modelValue"></div>
    </div>
  </div>
</template>

<script>
//https://quilljs.com/docs/api#getbounds
//https://www.kancloud.cn/liuwave/quill/1434140
import Quill from "quill";
import "quill/dist/quill.core.css";
import "quill/dist/quill.bubble.css";
import "quill/dist/quill.snow.css";
export default {
  name: "Editor",
  props: ["modelValue"],
  emits: ["update:modelValue"],
  data() {
    return {
      quill: null,
      editorConf: {
        readOnly: false,
        theme: "snow",
        placeholder: "",
        modules: {
          toolbar: {
            container: "",
            handlers: {},
          },
        },
      },
    };
  },
  mounted() {
    this.initEditor();
  },
  // beforeUnmount() {
  //   this.quill = null;
  //   delete this.quill;
  // },

  methods: {
    fileChange(e) {
      const _this = this;
      this.$uploadFile(e.target.files[0], (url) => {
        this.$nextTick(() => {
          _this.quill.insertEmbed(1, "image", url);
          console.log("插入图片", url, _this.quill.Target);
          // const range = _this.quill.getSelection();
          // if (range) {
          //   _this.quill.insertEmbed(range.index, "image", url);
          //   _this.quill.setSelection(range.index + 1);
          // }
        });
      });
    },
    //初始化编辑器
    initEditor() {
      this.editorConf.modules.toolbar.container = this.$refs.header;
      this.editorConf.modules.toolbar.handlers.image = () => {
        this.$refs.mfiles.click();
      };
      Quill.debug("info");
      this.quill = new Quill(this.$refs.editor_quill, this.editorConf);
      // this.quill.enable(false);
      // this.quill.setContents(this.modelValue);
      this.quill.enable(true);

      // this.quill.setContents(this.modelValue);

      // this.quill.on("text-change", () => {
      //   let html = this.$refs.Content.children[0].innerHTML;
      //   const text = this.quill.getText();
      //   if (html === "<p><br></p>") html = "";
      //   this.$emit("update:modelValue", html);
      // });
    },
  },
};
</script>

<style lang="scss" scoped>
.g-editor {
  border: thin solid gray;
  .header {
    text-align: left;
  }
  .body {
    text-align: left;
    width: 100%;
    min-height: 200px;
    img {
      max-width: 100%;
    }
  }
}
</style>
