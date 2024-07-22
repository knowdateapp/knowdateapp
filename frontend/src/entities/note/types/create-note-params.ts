import { Workspace } from 'shared/model';

export type CreateNoteParams = {
  workspace: Workspace;
  title?: string | null;
};
