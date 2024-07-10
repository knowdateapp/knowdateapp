import { Workspace } from 'shared/model';

export interface ICreateNoteParams {
  workspace: Workspace;
  title?: string | null;
}
