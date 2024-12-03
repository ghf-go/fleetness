<template>
  <div class="g-editor">
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

    <div class="body" ref="Content"></div>
    <input
      accept="image/*"
      ref="mfiles"
      type="file"
      style="display: none"
      @change="fileChange"
    />
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
  props: ["modelValue"],
  emits: ["update:modelValue"],
  data() {
    return {
      contentHtml: "",
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
  beforeDestroy() {
    this.quill = null;
    delete this.quill;
  },
  methods: {
      fileChange(f) {
        var _this = this
          this.$uploadFile(f, sucfunc(url){
            const range = _this.quill.getSelection();
            if (range) {
              _this.quill.insertEmbed(range.index, "image", url);
              _this.quill.setSelection(range.index + 1);
          })
    },
    //初始化编辑器
    initEditor() {
      this.editorConf.modules.toolbar.container = this.$refs.header;
      this.editorConf.modules.toolbar.handlers.image = () => {
        this.$refs.mfiles.click();
      };
      this.quill = new Quill(this.$refs.Content, this.editorConf);
      this.$refs.Content.children[0].innerHTML = this.modelValue;
      this.quill.enable(true);

      this.quill.on("text-change", () => {
        let html = this.$refs.Content.children[0].innerHTML;
        const quill = this.quill;
        const text = this.quill.getText();
        if (html === "<p><br></p>") html = "";
        this.$emit("update:modelValue", html);
      });
      this.$emit("ready", this.quill);
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
