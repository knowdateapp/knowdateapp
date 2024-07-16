import { apiClient } from 'shared/api';
import { Workspace } from 'shared/model';
import { ICreateNoteParams, INote } from '../types';

const getUrl = (workspace: Workspace) => `${workspace}/notes`;

export const controller = {
  createNote: ({ workspace, title = null }: ICreateNoteParams) =>
    apiClient.post<INote>(getUrl(workspace), { title }),
};
