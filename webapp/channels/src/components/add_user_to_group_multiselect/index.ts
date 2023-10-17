// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

import {connect} from 'react-redux';
import {bindActionCreators} from 'redux';
import type {ActionCreatorsMapObject, Dispatch} from 'redux';

import type {UserProfile} from '@mattermost/types/users';

import {getProfilesNotInGroup, searchProfiles, getProfiles} from 'mattermost-redux/actions/users';
import {getProfilesNotInCurrentGroup, getUserStatuses, getProfiles as getUsers} from 'mattermost-redux/selectors/entities/users';
import type {Action, ActionResult} from 'mattermost-redux/types/actions';

import {loadStatusesForProfilesList} from 'actions/status_actions';

import type {Value} from 'components/multiselect/multiselect';

import type {GlobalState} from 'types/store';

import AddUserToGroupMultiSelect from './add_user_to_group_multiselect';

type UserProfileValue = Value & UserProfile;

type OwnProps = {
    groupId?: string;
}

function mapStateToProps(state: GlobalState, props: OwnProps) {
    let profiles: UserProfileValue[];

    if (props.groupId) {
        profiles = getProfilesNotInCurrentGroup(state, props.groupId) as UserProfileValue[];
    } else {
        profiles = getUsers(state) as UserProfileValue[];
    }
    const userStatuses = getUserStatuses(state);

    return {
        groupId: props.groupId,
        profiles,
        userStatuses,
    };
}

type Actions = {
    getProfiles: (page?: number, perPage?: number) => Promise<ActionResult>;
    getProfilesNotInGroup: (groupId: string, page?: number, perPage?: number) => Promise<ActionResult>;
    loadStatusesForProfilesList: (users: UserProfile[]) => void;
    searchProfiles: (term: string, options: any) => Promise<ActionResult>;
}

function mapDispatchToProps(dispatch: Dispatch) {
    return {
        actions: bindActionCreators<ActionCreatorsMapObject<Action>, Actions>({
            getProfiles,
            getProfilesNotInGroup,
            loadStatusesForProfilesList,
            searchProfiles,
        }, dispatch),
    };
}

export default connect(mapStateToProps, mapDispatchToProps)(AddUserToGroupMultiSelect);
