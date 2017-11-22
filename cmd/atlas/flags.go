package main

/*
Flags management.

There are two types of flags for commands
- global (sort, page_size)
- local (dns-related ones, ping)related one s and so forth)

Outcome is the same in all cases, a given flag is transformed into a query flag
*/

var (
    globalFlags = map[string]string{
        "fields":          fFieldList,
        "format":          fFormat,
        "include":         fInclude,
        "optional_fields": fOptFields,
        "page":            fPageNum,
        "page_size":       fPageSize,
        "sort":            fSortOrder,
    }

    commonFlags = map[string]string{
    }
)

