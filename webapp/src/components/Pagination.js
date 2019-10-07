// Image Builder
// Copyright 2019 Canonical Ltd.  All rights reserved.

import React, { Component } from 'react'

class Pagination extends Component {

    constructor(props) {
        super(props)

        this.state = {
            //page: 1,
            query: null,
            maxRecords: props.pageSize | 10,
        }
    }

    pageUp = () => {
        let pages = this.calculatePages();
        let page = this.props.page + 1
        if (page > pages) {
            page = pages;
        }
        this.signalPageChange(page);
    }

    pageDown = () => {
        let page = this.props.page - 1
        if (page <= 0) {
            page = 1;
        }
        this.signalPageChange(page);
    }

    pageSet  = (e) => {
        let page = parseInt(e.target.getAttribute('data-key'))
        this.setState({page: page});
        this.signalPageChange(page);
    }

    signalPageChange(page) {
        // Signal the rows that the owner should display
        var startRow = ((page - 1) * this.state.maxRecords);

        this.props.pageChange(page, startRow, startRow + this.state.maxRecords);
    }

    calculatePages() {
        // Use the filtered row count when we a query has been entered
        var length = this.props.displayRows.length;

        var pages = parseInt(length / this.state.maxRecords, 10);
        if (length % this.state.maxRecords > 0) {
            pages += 1;
        }

        return pages;
    }

    renderPaging() {
        let pages = this.calculatePages();
        let band = parseInt((this.props.page - 0.1) / 3);
        if (pages > 1) {
            return (
                <div className="u-float--right spacer">
                    <button className="p-button--neutral small" onClick={this.pageDown} disabled={this.props.page===1 ? 'disabled': ''}>&lsaquo;</button>
                    <button className={((band * 3) + 1)===this.props.page ? 'p-button--neutral small active' : 'p-button--neutral small'} onClick={this.pageSet} data-key={(band * 3) + 1}>{(band * 3) + 1}</button>
                    {((band * 3) + 2) <= pages ? <button className={((band * 3) + 2)===this.props.page ? 'p-button--neutral small active' : 'p-button--neutral small'} onClick={this.pageSet} data-key={(band * 3) + 2}>{(band * 3) + 2}</button> : ''}
                    {((band * 3) + 3) <= pages ? <button className={((band * 3) + 3)===this.props.page ? 'p-button--neutral small active' : 'p-button--neutral small'} onClick={this.pageSet} data-key={(band * 3) + 3}>{(band * 3) + 3}</button> : ''}
                    <button className="p-button--neutral small" href="" onClick={this.pageUp} disabled={this.props.page===pages ? 'disabled': ''}>&rsaquo;</button>
                </div>
            );
        } else {
            return <div className="u-float--right" />;
        }
    }

    render() {
        return (
            <div className="col-12 pagination u-align--center">
                {this.renderPaging()}
            </div>
        );
    }
}

export default Pagination


//1 2 3 4 5 6 7 8 9 10
//--0--