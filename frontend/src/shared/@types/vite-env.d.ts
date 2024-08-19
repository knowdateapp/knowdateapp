/// <reference types="vite/client" />
import { Editor } from '@tiptap/react';

interface ImportMetaEnv {
  readonly VITE_API_URL: string;
}

interface ImportMeta {
  readonly env: ImportMetaEnv;
}

declare global {
  interface Window {
    editor: Editor | null;
  }
}
