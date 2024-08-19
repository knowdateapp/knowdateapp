import { Box } from '@chakra-ui/react';
import { EditorContent } from '@tiptap/react';
import { useRef } from 'react';
import { EditorHeader } from './components';
import { useTextEditor } from './hooks';
// import { LinkMenu } from '@/components/menus';

// import { Sidebar } from '@/components/Sidebar';
// import { EditorContext } from '@/context/EditorContext';
// import ImageBlockMenu from '@/extensions/ImageBlock/components/ImageBlockMenu';
// import { ColumnsMenu } from '@/extensions/MultiColumn/menus';
// import { TableColumnMenu, TableRowMenu } from '@/extensions/Table/menus';
// import { EditorHeader } from './components/EditorHeader';
// import { TextMenu } from '../menus/TextMenu';
// import { ContentItemMenu } from '../menus/ContentItemMenu';

export const TextEditor = () => {
  const menuContainerRef = useRef(null);
  const editorRef = useRef<HTMLDivElement | null>(null);

  const { editor, characterCount, leftSidebar } = useTextEditor('');

  // const providerValue = useMemo(() => {
  //   return {};
  // }, []);

  if (!editor) {
    return null;
  }

  return (
    // <EditorContext.Provider value={providerValue}>
    <Box display="flex" height="100%" ref={menuContainerRef}>
      {/*<Sidebar isOpen={leftSidebar.isOpen} onClose={leftSidebar.close} editor={editor} />*/}
      <Box position="relative" display="flex" flexDirection="column" flex={1} overflow="hidden">
        <EditorHeader
          characters={characterCount.characters()}
          words={characterCount.words()}
          isSidebarOpen={leftSidebar.isOpen}
          toggleSidebar={leftSidebar.toggle}
        />
        <EditorContent editor={editor} ref={editorRef} className="flex-1 overflow-y-auto" />
        {/*<ContentItemMenu editor={editor} />*/}
        {/*<LinkMenu editor={editor} appendTo={menuContainerRef} />*/}
        {/*<TextMenu editor={editor} />*/}
        {/*<ColumnsMenu editor={editor} appendTo={menuContainerRef} />*/}
        {/*<TableRowMenu editor={editor} appendTo={menuContainerRef} />*/}
        {/*<TableColumnMenu editor={editor} appendTo={menuContainerRef} />*/}
        {/*<ImageBlockMenu editor={editor} appendTo={menuContainerRef} />*/}
      </Box>
    </Box>
    // </EditorContext.Provider>
  );
};
