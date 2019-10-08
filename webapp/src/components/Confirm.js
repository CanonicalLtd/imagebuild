// Image Builder
// Copyright 2019 Canonical Ltd.  All rights reserved.

import React, {Component} from 'react';
import {T} from "./Utils";
import api from "./api";
import moment from "moment";

class Confirm extends Component {
    constructor(props) {
        super(props)
        this.state = {
            messages: [],
            buildURL: '',
            downloads: [],
            building: false,
        };
    }

    getBuildStatus() {
        let bld = {link: this.state.buildURL}
        let b = {}

        api.buildFetch(bld).then(response => {
            b = response.data.build
            let start = 'Build started: ' + b.date_started
            if (b.date_started) {
                start = start + ' (' + moment(b.date_started).fromNow() + ')'
            }
            let m = [start]
            m.push('Status: ' + b.buildstate)
            m.push('Type: ' + b.title)
            m.push('Log: ' + b.build_log_url)
            this.setState({messages: m, downloads: b.downloads})
        })
        .then( ()=> {
            if ((b.buildstate !== 'Successfully built') && (b.buildstate !== 'Failed to build')) {
                this.poll()
            } else {
                this.setState({building: false})
            }
        })
    }

    poll = () => {
        // Poll every second
        setTimeout(this.getBuildStatus.bind(this), 10000);
    }

    handleBuild = (e) => {
        let selected = this.props.board
        let bld = {boardId: selected.board.id, osId: selected.os.id, snaps: selected.snaps}

        // Start the build
        api.buildRequest(bld).then(response => {
            this.setState({buildURL: response.data.buildURL, building: true})
        }).then( ()=> {
            // Monitor the build status
            this.getBuildStatus()
        })
    }

    renderSelected() {
        return (
            <div className="row summary">
                <p><i className="p-icon--success" />{this.props.board.board.name}</p>
                <p><i className="p-icon--success" />Ubuntu {T(this.props.board.os.type)} {this.props.board.os.version}</p>
                <p><i className="p-icon--success" />{this.props.board.snaps.length !== 1 ? this.props.board.snaps.length + ' ' + T('snaps'): '1 '+ T('snap')}</p>
            </div>
        )
    }

    renderConsole() {
        if (this.state.messages.length === 0) {
            return
        }
        return (
            <pre className="console">
                {this.state.messages.map(m => {
                    return m + '\n'
                })}
            </pre>
        )
    }

    renderDownloads() {
        if ((!this.state.downloads) || (this.state.downloads.length === 0)) {
            return
        }
        return (
            <div className="row">
                <h2>{T('downloads')}</h2>
                {this.state.downloads.map(d => {
                    let parts = d.split('/')
                    let filename = parts[parts.length-1]
                    console.log('---', filename)
                    let extensions = filename.split('.')
                    let suffix = extensions[extensions.length-1]
                    if (suffix==='xz') {
                        return <p><a href={d}>{filename}</a></p>
                    }
                })}
            </div>
        )
    }

    renderMessage() {
        if (this.state.building) {
            return (
                <div>
                    <img src="/static/images/ajax-loader.gif" /> {T('build-in-progress')}
                </div>
            )
        } else {
            return (
                <div>
                    {T('ready-to-build')} &nbsp;
                    <button className="p-button--brand" onClick={this.handleBuild} disabled={this.state.building}>{T('build')}</button>
                </div>
            )
        }
    }

    render() {
        let selected = this.props.board
        return (
            <div className="first">
                <h2>{T('confirm-configuration')}</h2>
                {this.renderSelected()}

                <div className="row">
                    <pre>
                        <code>
                            board:{'\n'}
                            {}  name: {selected.board.name}{'\n'}
                            {}  os:{'\n'}
                            {}    type: {selected.os.type}{'\n'}
                            {}    version: {selected.os.version}{'\n'}
                            {}  snaps:{'\n'}
                            {selected.snaps.map(s => {
                                return '    ' + s + '\n'
                            })}
                        </code>
                    </pre>
                </div>

                {this.renderMessage()}

                <div className="row">
                    {this.renderConsole()}
                </div>
                {this.renderDownloads()}
            </div>
        );
    }
}

export default Confirm;