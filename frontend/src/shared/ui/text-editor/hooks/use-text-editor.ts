import { useEditor } from '@tiptap/react';
// import { ExtensionKit } from '../../ui/text-editor/components/extension-kit';
import { StarterKit } from '@tiptap/starter-kit';
import { useSidebar } from './use-sidebar.ts';

export const useTextEditor = (initialContent: string) => {
  const leftSidebar = useSidebar();

  const editor = useEditor(
    {
      autofocus: true,
      onCreate: ({ editor }) => {
        if (editor.isEmpty) {
          editor.commands.setContent(initialContent);
        }
      },
      // extensions: [...ExtensionKit()],
      extensions: [StarterKit],
      editorProps: {
        attributes: {
          autocomplete: 'off',
          autocorrect: 'off',
          autocapitalize: 'off',
          class: 'min-h-full',
        },
      },
    },
    [initialContent],
  );

  // TODO уебаны не типизировали, поправить
  const characterCount = (editor?.storage.characterCount as Record<string, unknown>) ?? {
    characters: () => 0,
    words: () => 0,
  };

  window.editor = editor;

  return { editor, characterCount, leftSidebar };
};
