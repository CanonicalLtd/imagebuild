// Image Builder
// Copyright 2019 Canonical Ltd.  All rights reserved.

import React, {Component} from 'react';
import api from './api';
import Pagination from './Pagination';
import Constants from './constants';
import {T} from './Utils';
import filesize from 'filesize'

const PAGELENGTH = 10;

class SnapDialogBox extends Component {

    constructor(props) {
        super(props)
        this.state = {
            currentSearch: '',
            snapName: '',
            snaps: [],
            loadingSearch: false,

            page: 1,
            startRow: 0,
            endRow: PAGELENGTH,
        };
    }

    searchStore() {
        if (this.state.snapName.length===0) {
            this.setState({currentSearch: ''})
            return
        }
        let bld = {boardId: this.props.board.id, osId: this.props.os.id}

        this.setState({loadingSearch: true})

        api.storeSearch(this.state.snapName, bld).then(response => {
            if ((response.data._embedded) && (response.data._embedded['clickindex:package'])) {
                this.setState({snaps: response.data._embedded['clickindex:package'], loadingSearch: false, message: null, messageType: null, currentSearch: this.state.snapName, page: 1, startRow: 0, endRow: PAGELENGTH})
            }
        })
    }

    handleClear = (e) => {
        this.setState({snaps: [], currentSearch: ''})
    }

    handleSearchChange = (e) => {
        e.preventDefault()
        this.setState({snapName: e.target.value})
    }

    handleKeyPress = (e) => {
        if (e.key === 'Enter') {
            e.stopPropagation()
            this.searchStore()
        }
    }

    handleSearchStore = (e) => {
        e.preventDefault()
        this.searchStore()
    }

    handleInstall = (e) => {
        e.preventDefault()
        var snap = e.target.getAttribute('data-key')

        this.props.handleInstallClick(snap);
    }

    handleRecordsForPage = (page, startRow, endRow) => {
        this.setState({page: page, startRow: startRow, endRow: endRow});
    }

    renderSnaps(snaps) {
        if (snaps.length > 0) {

            return (
                <div>
                    <p>{snaps.length} snaps found</p>
                    <table className="store">
                        <thead>
                            <tr>
                                <th>{T('snap-name')}</th>
                            </tr>
                        </thead>
                        <tbody>
                        {snaps.slice(this.state.startRow, this.state.endRow).map(s => {
                            return (
                                <tr key={s.snap_id} title={s.description}>
                                    <td className="overflow">
                                        <img src={s.icon_url || Constants.missingIcon} width="30px" />
                                    
                                        <button data-key={s.package_name} className="p-button--neutral small" title={T("add-application")} onClick={this.handleInstall}>
                                            {T('add')}
                                        </button>

                                        <b>{s.package_name}</b> {s.version}<br />
                                        {s.developer_name} ({filesize(s.binary_filesize)})
                                    </td>
                                </tr>
                            )
                        })}
                        </tbody>
                    </table>
                </div>
            )
        } else if (this.state.currentSearch.length > 0) {
            return <div>No snaps found.</div>
        }
    }

    render() {
        var snaps = this.state.snaps

        if (this.props.message) {
            return (
                <div className="col-6 snaps">
                    <h4>
                        {this.state.loadingSearch ? <img src={Constants.LoadingImage} alt={T('loading')} /> : ''}
                        {this.props.message}
                    </h4>
                    <p>
                        <form className="p-search-box">
                            <input className="p-search-box__input" type="search" name="snapname" onKeyPress={this.handleKeyPress} onChange={this.handleSearchChange} placeholder={T('search-store')} />
                            <button type="reset" className="p-search-box__reset" alt="reset" disabled="" onClick={this.handleClear}><i className="p-icon--close" /></button>
                            <button type="submit" onClick={this.handleSearchStore} className="p-search-box__button" alt="search"><i className="p-icon--search" /></button>
                        </form>
                    </p>

                    {this.renderSnaps(snaps)}

                    <Pagination displayRows={snaps}
                                pageSize={PAGELENGTH}
                                page={this.state.page}
                                pageChange={this.handleRecordsForPage} />
                </div>
            );
        } else {
            return <span />;
        }
    }
}

export default SnapDialogBox;