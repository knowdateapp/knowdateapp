import { apiClient } from 'shared/api';
import { Workspace } from 'shared/model';
import { CreateNoteParams, NoteEntity } from '../types';

const getUrl = (workspace: Workspace) => `${workspace}/notes`;

export const controller = {
  createNote: ({ workspace, title = null }: CreateNoteParams) =>
    apiClient.post<NoteEntity>(getUrl(workspace), { title }),
};
