import { Workspace } from 'shared/types';

export type CreateNoteParams = {
  workspace: Workspace;
  title?: string | null;
};
